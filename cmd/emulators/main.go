package main

import (
	"context"
	_ "embed"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/subscribeddotdev/subscribed-backend/tests/jwks"
	"golang.org/x/sync/errgroup"
)

const (
	JwksServerPort    = 8090
	WebhookServerPort = 8091
)

type GeneralResponse struct {
	Message string `json:"message"`
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	g, _ := errgroup.WithContext(context.Background())

	g.Go(func() error {
		return runWebhookServer(logger, WebhookServerPort)
	})

	g.Go(func() error {
		return runJwksServer(logger, JwksServerPort)
	})

	err := g.Wait()
	if err != nil {
		logger.Error("emulators.. service crashed", "error", err)
		os.Exit(1)
	}
}

// This server emulates webhook endpoints that would have been provided by
// customers
func runWebhookServer(logger *slog.Logger, port int) error {
	router := echo.New()
	commonMiddlewares(router)

	router.POST("/webhook", func(c echo.Context) error {
		// Simulate latency
		time.Sleep(time.Millisecond * time.Duration(gofakeit.Number(80, 500)))

		action := c.Param("action")
		if action == "invalid-credentials" {
			return c.JSON(http.StatusUnauthorized, GeneralResponse{Message: "invalid-credentials"})
		}

		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		return c.String(http.StatusOK, fmt.Sprintf(`{"keys": [%s]}`, jwks.JwksPublicKey))
	})

	logger.Info("[Webhook Emulator] server about to start at", "port", port)
	err := router.Start(fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("webhook emulator crashed: %v", err)
	}

	return nil
}

func runJwksServer(logger *slog.Logger, port int) error {
	router := echo.New()
	commonMiddlewares(router)

	router.GET("/jwks", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		return c.String(http.StatusOK, fmt.Sprintf(`{"keys": [%s]}`, jwks.JwksPublicKey))
	})

	logger.Info("[JWKS Emulator] server about to start at", "port", port)
	err := router.Start(fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("jwks emulator crashed: %v", err)
	}

	return nil
}

func commonMiddlewares(router *echo.Echo) {
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
}
