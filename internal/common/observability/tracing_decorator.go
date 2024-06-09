package observability

import (
	"context"
)

type commandTracingDecorator[C any] struct {
	// tracer trace.Tracer
	base CommandHandler[C]
}

type commandWithResultTracingDecorator[C any, R any] struct {
	// tracer trace.Tracer
	base CommandWithResultHandler[C, R]
}

type queryTracingDecorator[Q any, R any] struct {
	// tracer trace.Tracer
	base QueryHandler[Q, R]
}

func (c commandTracingDecorator[C]) Execute(ctx context.Context, cmd C) error {
	// ctx, span := c.tracer.Start(ctx, handlerName(c.base))
	// defer span.End()

	return c.base.Execute(ctx, cmd)
}

func (c commandWithResultTracingDecorator[C, R]) Execute(ctx context.Context, cmd C) (R, error) {
	// ctx, span := c.tracer.Start(ctx, handlerName(c.base))
	// defer span.End()
	return c.base.Execute(ctx, cmd)
}

func (c queryTracingDecorator[Q, R]) Execute(ctx context.Context, q Q) (R, error) {
	// ctx, span := c.tracer.Start(ctx, handlerName(c.base))
	// defer span.End()

	return c.base.Execute(ctx, q)
}
