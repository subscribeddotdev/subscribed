package messaging

import (
	"encoding/json"
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/oklog/ulid/v2"
)

type Publisher struct {
	wmPublisher message.Publisher
}

func NewPublisher(config amqp.Config, logger watermill.LoggerAdapter) (*Publisher, error) {
	publisher, err := amqp.NewPublisher(config, logger)
	if err != nil {
		return nil, err
	}

	return &Publisher{
		wmPublisher: publisher,
	}, nil
}

func (p *Publisher) Publish(topic string, event any) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("unable to marshall event from '%s' due to: %v", topic, err)
	}

	msg := message.NewMessage(ulid.Make().String(), payload)

	return p.wmPublisher.Publish(topic, msg)
}

func (p *Publisher) Close() error {
	return p.wmPublisher.Close()
}
