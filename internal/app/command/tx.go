package command

import (
	"context"
)

type TransactFunc func(adapters TransactableAdapters) error

type TransactionProvider interface {
	Transact(ctx context.Context, f TransactFunc) error
}

type TransactableAdapters struct {
	EventPublisher EventPublisher
}

type EventPublisher interface {
}
