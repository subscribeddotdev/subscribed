package app

import (
	"context"

	"github.com/subscribeddotdev/subscribed/server/internal/app/auth"
	"github.com/subscribeddotdev/subscribed/server/internal/app/command"
	"github.com/subscribeddotdev/subscribed/server/internal/app/query"
	"github.com/subscribeddotdev/subscribed/server/internal/domain"
	"github.com/subscribeddotdev/subscribed/server/internal/domain/iam"
)

type CommandHandler[C any] interface {
	Execute(ctx context.Context, cmd C) error
}

// CommandHandlerWithResult To be used in rare occasions.
type CommandHandlerWithResult[C, R any] interface {
	Execute(ctx context.Context, cmd C) (result R, err error)
}

type QueryHandler[Q, R any] interface {
	Execute(ctx context.Context, q Q) (R, error)
}

type Command struct {
	AddEndpoint         CommandHandler[command.AddEndpoint]
	CallWebhookEndpoint CommandHandler[command.CallWebhookEndpoint]
	SendMessage         CommandHandler[command.SendMessage]

	// Applications
	CreateApplication CommandHandlerWithResult[command.CreateApplication, domain.ApplicationID]

	// IAM
	SignUp CommandHandler[command.Signup]
	SignIn CommandHandlerWithResult[command.SignIn, *iam.Member]

	// Event types
	CreateEventType CommandHandlerWithResult[command.CreateEventType, domain.EventTypeID]

	// Api keys
	CreateApiKey  CommandHandlerWithResult[command.CreateApiKey, *domain.ApiKey]
	DestroyApiKey CommandHandler[command.DestroyApiKey]
}

type Query struct {
	AllApiKeys      QueryHandler[query.AllApiKeys, []*domain.ApiKey]
	AllEnvironments QueryHandler[query.AllEnvironments, []*domain.Environment]

	// Applications
	Application     QueryHandler[query.Application, *domain.Application]
	AllApplications QueryHandler[query.AllApplications, query.Paginated[[]domain.Application]]
}

type App struct {
	Authorization *auth.Service
	Command       Command
	Query         Query
}
