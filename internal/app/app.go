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
	AddEndpoint        CommandHandler[command.AddEndpoint]
	SendMessage        CommandHandler[command.SendMessage]
	CreateApplication  CommandHandler[command.CreateApplication]
	CreateOrganization CommandHandler[command.CreateOrganization]
}

type App struct {
	Command Command
}
