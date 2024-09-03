package messaging

import (
	"encoding/json"
	"fmt"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/oklog/ulid/v2"
	"github.com/subscribeddotdev/subscribed/server/internal/common/logs"
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

type AmqpPublisher struct {
	publisher message.Publisher
}

func NewAmqpPublisher(amqpUrl string, logger *logs.Logger) (AmqpPublisher, error) {
	amqpPublisher, err := amqp.NewPublisher(amqp.NewDurableQueueConfig(amqpUrl), watermill.NewSlogLogger(logger.Logger))
	if err != nil {
		return AmqpPublisher{}, err
	}

	return AmqpPublisher{
		publisher: NewCorrelationIdPublisher(amqpPublisher),
	}, nil
}

func (p AmqpPublisher) Publish(topic string, messages ...*message.Message) error {
	return p.publisher.Publish(topic, messages...)
}

func (p AmqpPublisher) Close() error {
	return p.publisher.Close()
}
