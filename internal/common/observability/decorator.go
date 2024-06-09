package observability

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

type QueryHandler[Q any, R any] interface {
	Execute(ctx context.Context, q Q) (result R, err error)
}

type CommandHandler[C any] interface {
	Execute(ctx context.Context, cmd C) (err error)
}

type CommandWithResultHandler[C any, R any] interface {
	Execute(ctx context.Context, cmd C) (result R, err error)
}

func handlerName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}

func prettyPrint(v any) string {
	value, err := json.Marshal(v)
	if err != nil {
		return "<marshalling error>"
	}

	return string(value)
}
