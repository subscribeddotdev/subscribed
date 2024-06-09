package app

import (
	"context"
)

type CommandHandler[C any] interface {
	Execute(ctx context.Context, cmd C) error
}

type QueryHandler[Q, R any] interface {
	Execute(ctx context.Context, q Q) error
}

type Command struct {
}

type App struct {
	Command Command
}
