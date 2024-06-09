package observability

import "context"

type commandMetricsDecorator[C any] struct {
	base CommandHandler[C]
}

type commandWithResultMetricsDecorator[C any, R any] struct {
	base CommandWithResultHandler[C, R]
}

type queryMetricsDecorator[Q any, R any] struct {
	base QueryHandler[Q, R]
}

func (c commandMetricsDecorator[C]) Execute(ctx context.Context, cmd C) error {
	return c.base.Execute(ctx, cmd)
}

func (c commandWithResultMetricsDecorator[C, R]) Execute(ctx context.Context, cmd C) (R, error) {
	return c.base.Execute(ctx, cmd)
}

func (c queryMetricsDecorator[Q, R]) Execute(ctx context.Context, q Q) (R, error) {
	return c.base.Execute(ctx, q)
}
