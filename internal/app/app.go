package app

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/app/auth"
	"github.com/subscribeddotdev/subscribed-backend/internal/app/command"
	"github.com/subscribeddotdev/subscribed-backend/internal/app/query"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
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
	SendMessage         CommandHandler[command.SendMessage]
	CreateApplication   CommandHandler[command.CreateApplication]
	SignUp              CommandHandler[command.Signup]
	SignIn              CommandHandlerWithResult[command.SignIn, *iam.Member]
	CreateEventType     CommandHandler[command.CreateEventType]
	CreateApiKey        CommandHandlerWithResult[command.CreateApiKey, *domain.ApiKey]
	CallWebhookEndpoint CommandHandler[command.CallWebhookEndpoint]
}

type Query struct {
	Environments QueryHandler[query.Environments, []*domain.Environment]
	AllApiKeys   QueryHandler[query.AllApiKeys, []*domain.ApiKey]
}

type App struct {
	Authorization *auth.Service
	Command       Command
	Query         Query
}
