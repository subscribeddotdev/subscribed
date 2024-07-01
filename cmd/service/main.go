package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/psql"
	"github.com/subscribeddotdev/subscribed-backend/internal/app/auth"
	"github.com/subscribeddotdev/subscribed-backend/internal/app/command"
	"github.com/subscribeddotdev/subscribed-backend/internal/common/clerkhttp"
	"github.com/subscribeddotdev/subscribed-backend/internal/common/observability"
	svix "github.com/svix/svix-webhooks/go"
	"golang.org/x/sync/errgroup"

	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/events"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/transaction"
	"github.com/subscribeddotdev/subscribed-backend/internal/app"
	"github.com/subscribeddotdev/subscribed-backend/internal/common/logs"
	"github.com/subscribeddotdev/subscribed-backend/internal/common/postgres"
	"github.com/subscribeddotdev/subscribed-backend/internal/ports/http"
)

type Config struct {
	DatabaseUrl            string `envconfig:"DATABASE_URL" required:"true"`
	Port                   int    `envconfig:"HTTP_PORT" required:"true"`
	ProductionMode         bool   `envconfig:"PRODUCTION_MODE" required:"true"`
	AllowedCorsOrigin      string `envconfig:"HTTP_ALLOWED_CORS" required:"true"`
	AmqpUrL                string `envconfig:"AMQP_URL" required:"true"`
	ClerkSecretKey         string `envconfig:"CLERK_SECRET_KEY" required:"true"`
	ClerkEmulatorServerURL string `envconfig:"CLERK_EMULATOR_SERVER_URL" required:"true"`
	ClerkWebhookSecret     string `envconfig:"CLERK_WEBHOOK_SECRET" required:"true"`
}

func main() {
	logger := logs.New()

	if err := run(logger); err != nil {
		logger.Fatal("service crashed due to an error", "error", err)
	}

	os.Exit(0)
}

func run(logger *logs.Logger) error {
	config := &Config{}
	err := envconfig.Process("", config)
	if err != nil {
		return fmt.Errorf("unable to load env variables: %v", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	g, gctx := errgroup.WithContext(ctx)

	db, err := postgres.Connect(config.DatabaseUrl)
	if err != nil {
		return err
	}

	if !config.ProductionMode {
		logger.Info("[LOGIN PROVIDER] setting up clerk to work with the simulator", "emulator_url", config.ClerkEmulatorServerURL)
		clerkhttp.SetupClerkForTestingMode(config.ClerkEmulatorServerURL)
	}

	applicationRepo := psql.NewApplicationRepository(db)
	endpointRepo := psql.NewEndpointRepository(db)
	eventTypeRepo := psql.NewEventTypeRepository(db)
	memberRepo := psql.NewMemberRepository(db)
	apiKeyRepo := psql.NewApiKeyRepository(db)
	envRepo := psql.NewEnvironmentRepository(db)

	eventPublisher, err := events.NewPublisher(config.AmqpUrL, watermill.NewStdLogger(!config.ProductionMode, !config.ProductionMode))
	if err != nil {
		return err
	}

	txProvider := transaction.NewPsqlProvider(db, eventPublisher, logger)

	application := &app.App{
		Authorization: auth.NewService(memberRepo),
		Command: app.Command{
			CreateOrganization: observability.NewCommandDecorator[command.CreateOrganization](command.NewCreateOrganizationHandler(txProvider), logger),
			CreateApplication:  observability.NewCommandDecorator[command.CreateApplication](command.NewCreateApplicationHandler(applicationRepo), logger),
			AddEndpoint:        observability.NewCommandDecorator[command.AddEndpoint](command.NewAddEndpointHandler(endpointRepo), logger),
			SendMessage:        observability.NewCommandDecorator[command.SendMessage](command.NewSendMessageHandler(txProvider), logger),
			CreateEventType:    observability.NewCommandDecorator[command.CreateEventType](command.NewCreateEventTypeHandler(eventTypeRepo), logger),
			CreateApiKey:       observability.NewCommandDecorator[command.CreateApiKey](command.NewCreateApiKeyHandler(apiKeyRepo, envRepo), logger),
		},
	}

	var webhookVerifier http.LoginProviderWebhookVerifier
	if config.ProductionMode {
		webhookVerifier, err = svix.NewWebhook(config.ClerkWebhookSecret)
		if err != nil {
			return fmt.Errorf("unable to create a webhook verifier: %v", err)
		}
	} else {
		webhookVerifier = &clerkhttp.MockWebHookVerifier{}
	}

	httpserver, err := http.NewServer(http.Config{
		Ctx:                          ctx,
		Logger:                       logger,
		Application:                  application,
		Port:                         config.Port,
		IsDebug:                      !config.ProductionMode,
		ClerkSecretKey:               config.ClerkSecretKey,
		LoginProviderWebhookVerifier: webhookVerifier,
		AllowedCorsOrigin:            strings.Split(config.AllowedCorsOrigin, ","),
	})
	if err != nil {
		return err
	}

	g.Go(func() error {
		return httpserver.Start()
	})

	// Gracefully termination of services
	g.Go(func() error {
		<-gctx.Done()

		logger.Info("starting gracefully termination")

		tCtx, tCancel := context.WithTimeout(context.Background(), time.Second*5)
		defer tCancel()

		err = httpserver.Stop(tCtx)
		if err != nil {
			return err
		}

		logger.Info("service terminated gracefully")

		return nil
	})

	return g.Wait()
}
