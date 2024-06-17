package app

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/app/command"
)

type CommandHandler[C any] interface {
	Execute(ctx context.Context, cmd C) error
}

type QueryHandler[Q, R any] interface {
	Execute(ctx context.Context, q Q) error
}

type Command struct {
	// Application
	CreateApplication CommandHandler[command.CreateApplication]

	// Iam
	CreateOrganization CommandHandler[command.CreateOrganization]
}

type App struct {
	Command Command
}
