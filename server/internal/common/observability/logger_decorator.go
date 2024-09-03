package observability

import (
	"context"
	"log/slog"

	"github.com/subscribeddotdev/subscribed/server/internal/common/logs"
)

type commandLoggingDecorator[C any] struct {
	logger *logs.Logger
	base   CommandHandler[C]
}

func (c commandLoggingDecorator[C]) Execute(ctx context.Context, cmd C) (err error) {
	logger := c.logger.With(
		slog.String("name", handlerName(c.base)),
		slog.String("command", prettyPrint(cmd)),
	)

	logger.Debug("Executing command")
	defer func() {
		if err == nil {
			logger.Info("Command executed successfully")
		} else {
			logger.Error("Failed to execute command", "error", err)
		}
	}()

	return c.base.Execute(ctx, cmd)
}

type commandWithResultLoggingDecorator[Q any, R any] struct {
	logger *logs.Logger
	base   QueryHandler[Q, R]
}

func (c commandWithResultLoggingDecorator[C, R]) Execute(ctx context.Context, cmd C) (result R, err error) {
	logger := c.logger.With(
		slog.String("name", handlerName(c.base)),
		slog.String("command_with_result", prettyPrint(cmd)),
	)

	logger.Debug("Executing command with result")
	defer func() {
		if err == nil {
			logger.Info("Command with result executed successfully")
		} else {
			logger.Error("Failed to execute command with result", "error", err)
		}
	}()

	return c.base.Execute(ctx, cmd)
}

type queryLoggingDecorator[Q any, R any] struct {
	logger *logs.Logger
	base   QueryHandler[Q, R]
}

func (c queryLoggingDecorator[Q, R]) Execute(ctx context.Context, q Q) (result R, err error) {
	logger := c.logger.With(
		slog.String("name", handlerName(c.base)),
		slog.String("command", prettyPrint(q)),
	)

	logger.Debug("Executing query")
	defer func() {
		if err == nil {
			logger.Info("Query executed successfully")
		} else {
			logger.Error("Failed to execute query", "error", err)
		}
	}()

	return c.base.Execute(ctx, q)
}
