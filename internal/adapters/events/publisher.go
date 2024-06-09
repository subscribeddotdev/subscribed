package events

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"

	"github.com/subscribeddotdev/subscribed-backend/internal/common/messaging"
)

type Publisher struct {
	eventPublisher *messaging.Publisher
}

func NewPublisher(amqpUrl string, logger watermill.LoggerAdapter) (*Publisher, error) {
	eventPublisher, err := messaging.NewPublisher(amqp.NewDurableQueueConfig(amqpUrl), logger)
	if err != nil {
		return nil, err
	}

	return &Publisher{
		eventPublisher: eventPublisher,
	}, nil
}
