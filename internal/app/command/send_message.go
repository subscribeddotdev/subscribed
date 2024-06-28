package command

import (
	"context"
	"fmt"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type SendMessage struct {
	EventTypeID   string
	Payload       string
	ApplicationID string
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
	eventTypeID, err := domain.NewIdFromString(cmd.EventTypeID)
	if err != nil {
		return err
	}

	applicationID, err := domain.NewIdFromString(cmd.ApplicationID)
	if err != nil {
		return err
	}

	message, err := domain.NewMessage(eventTypeID, applicationID, cmd.Payload)
	if err != nil {
		return err
	}

	return c.txProvider.Transact(ctx, func(adapters TransactableAdapters) error {
		err = adapters.MessageRepository.Insert(ctx, message)
		if err != nil {
			return fmt.Errorf("error saving message: %v", err)
		}

		return nil
	})
}
