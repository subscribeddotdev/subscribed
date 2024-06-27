package command

import (
	"context"
)

type SendMessage struct {
}

type SendMessageHandler struct {
	txProvider TransactionProvider
}

func NewSendMessageHandler(txProvider TransactionProvider) SendMessageHandler {
	return SendMessageHandler{
		txProvider: txProvider,
	}
}

func (c SendMessageHandler) Execute(ctx context.Context, cmd SendMessage) error {
	return c.txProvider.Transact(ctx, func(adapters TransactableAdapters) error {
		return nil
	})
}
