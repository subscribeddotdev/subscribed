package command

import (
	"context"
	"fmt"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type SendMessage struct {
	EventTypeID   domain.EventTypeID
	Payload       string
	ApplicationID domain.ApplicationID
	OrgID         string
}

type SendMessageHandler struct {
	txProvider   TransactionProvider
	endpointRepo domain.EndpointRepository
}

func NewSendMessageHandler(txProvider TransactionProvider, endpointRepo domain.EndpointRepository) SendMessageHandler {
	return SendMessageHandler{
		txProvider:   txProvider,
		endpointRepo: endpointRepo,
	}
}

func (c SendMessageHandler) Execute(ctx context.Context, cmd SendMessage) error {
	message, err := domain.NewMessage(
		cmd.EventTypeID,
		cmd.OrgID,
		cmd.ApplicationID,
		cmd.Payload,
	)
	if err != nil {
		return err
	}

	// TODO: consider moving this inside the transaction???
	endpoints, err := c.endpointRepo.ByEventTypeIdAndAppID(ctx, message.EventTypeID(), message.ApplicationID())
	if err != nil {
		return fmt.Errorf("error retrieving endpoints: %v", err)
	}

	return c.txProvider.Transact(ctx, func(adapters TransactableAdapters) error {
		err = adapters.MessageRepository.Insert(ctx, message)
		if err != nil {
			return fmt.Errorf("error saving message: %v", err)
		}

		// TODO: test that a message is published for each endpoint subscribed to the event_type_id
		for _, endpoint := range endpoints {
			err = adapters.EventPublisher.PublishMessageSent(ctx, MessageSent{
				MessageID:  message.Id().String(),
				EndpointID: endpoint.ID().String(),
			})
			if err != nil {
				return fmt.Errorf("error publishing the event MessageSent: %v", err)
			}
		}

		return nil
	})
}
