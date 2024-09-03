package observability

import (
	"context"

	"github.com/subscribeddotdev/subscribed/server/internal/common/logs"
)

type commandDecorator[C any] struct {
	base CommandHandler[C]
}

func NewCommandDecorator[C any](base CommandHandler[C], logger *logs.Logger) commandDecorator[C] {
	return commandDecorator[C]{
		base: commandMetricsDecorator[C]{
			base: commandTracingDecorator[C]{
				base: commandLoggingDecorator[C]{
					base:   base,
					logger: logger,
				},
			},
		},
	}
}

func (q commandDecorator[C]) Execute(ctx context.Context, cmd C) error {
	return q.base.Execute(ctx, cmd)
}
