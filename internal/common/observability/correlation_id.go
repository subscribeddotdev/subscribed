package observability

import (
	"context"
	"errors"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type contextKey string

const correlationIdContextKey contextKey = "correlation_id"

func NewCorrelationID() string {
	return domain.NewID().String()
}

func ContextWithCorrelationID(ctx context.Context, correlationID string) context.Context {
	newCtx := context.WithValue(ctx, correlationIdContextKey, contextKey(correlationID))
	return newCtx
}

func CorrelationIdFromContext(ctx context.Context) (string, error) {
	corrID, ok := ctx.Value(correlationIdContextKey).(contextKey)
	if !ok {
		return "", errors.New("correlation ID not found in context")
	}

	return string(corrID), nil
}
