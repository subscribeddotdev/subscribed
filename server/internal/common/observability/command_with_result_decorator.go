package observability

import (
	"context"

	"github.com/subscribeddotdev/subscribed/server/internal/common/logs"
)

type commandWithResultDecorator[C, R any] struct {
	base CommandWithResultHandler[C, R]
}

func NewCommandWithResultDecorator[C, R any](base CommandWithResultHandler[C, R], logger *logs.Logger) commandWithResultDecorator[C, R] {
	return commandWithResultDecorator[C, R]{
		base: commandWithResultMetricsDecorator[C, R]{
			base: commandWithResultTracingDecorator[C, R]{
				base: commandWithResultLoggingDecorator[C, R]{
					base:   base,
					logger: logger,
				},
			},
		},
	}
}

func (q commandWithResultDecorator[C, R]) Execute(ctx context.Context, cmd C) (result R, err error) {
	return q.base.Execute(ctx, cmd)
}
