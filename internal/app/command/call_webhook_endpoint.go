package command

import (
	"context"
	"fmt"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type CallWebhookEndpoint struct {
	EndpointID string
	MessageID  string
}

type CallWebhookEndpointHandler struct {
	txProvider TransactionProvider
}

func NewCallWebhookEndpointHandler(txProvider TransactionProvider) CallWebhookEndpointHandler {
	return CallWebhookEndpointHandler{
		txProvider: txProvider,
	}
}

func (c CallWebhookEndpointHandler) Execute(ctx context.Context, cmd CallWebhookEndpoint) error {
	endpointID, err := domain.NewIdFromString(cmd.EndpointID)
	if err != nil {
		return err
	}

	messageID, err := domain.NewIdFromString(cmd.MessageID)
	if err != nil {
		return err
	}

	return c.txProvider.Transact(ctx, func(adapters TransactableAdapters) error {
		endpoint, err := adapters.EndpointRepository.ByID(ctx, endpointID)
		if err != nil {
			return fmt.Errorf("error querying the endpoint by id '%s': %v", cmd.EndpointID, err)
		}

		message, err := adapters.MessageRepository.ByID(ctx, messageID)
		if err != nil {
			return fmt.Errorf("error querying message by id '%s': %v", cmd.MessageID, err)
		}

		fmt.Printf("\n\n")
		fmt.Println(message.OrgID(), message.Payload(), endpoint.EndpointURL(), endpoint.SigningSecret())
		fmt.Printf("\n\n")
		// Load message
		// Load endpoint
		// Transform the message using the schema
		// Call the endpoint
		// Record the call result
		return nil
	})
}
