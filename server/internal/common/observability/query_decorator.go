package observability

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/common/logs"
)

type queryDecorator[Q any, R any] struct {
	logger *logs.Logger
	base   QueryHandler[Q, R]
}

func NewQueryDecorator[Q any, R any](base QueryHandler[Q, R], logger *logs.Logger) queryDecorator[Q, R] {
	return queryDecorator[Q, R]{
		logger: logger,
		base: queryMetricsDecorator[Q, R]{
			base: queryTracingDecorator[Q, R]{
				base: queryLoggingDecorator[Q, R]{
					base:   base,
					logger: logger,
				},
			},
		},
	}
}

func (q queryDecorator[Q, R]) Execute(ctx context.Context, query Q) (result R, err error) {
	return q.base.Execute(ctx, query)
}
