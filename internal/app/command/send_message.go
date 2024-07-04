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
	orgID, err := domain.NewIdFromString(cmd.OrgID)
	if err != nil {
		return err
	}

	eventTypeID, err := domain.NewIdFromString(cmd.EventTypeID)
	if err != nil {
		return err
	}

	applicationID, err := domain.NewIdFromString(cmd.ApplicationID)
	if err != nil {
		return err
	}

	message, err := domain.NewMessage(eventTypeID, orgID, applicationID, cmd.Payload)
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

		for _, endpoint := range endpoints {
			err = adapters.EventPublisher.PublishMessageSent(ctx, MessageSent{
				MessageID:  message.Id().String(),
				EndpointID: endpoint.Id().String(),
			})
			if err != nil {
				return fmt.Errorf("error publishing the event MessageSent: %v", err)
			}
		}

		return nil
	})
}
