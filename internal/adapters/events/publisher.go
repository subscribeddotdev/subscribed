package events

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"
	"github.com/ThreeDotsLabs/watermill/message"
	events "github.com/subscribeddotdev/subscribed-backend/events/go"
	"github.com/subscribeddotdev/subscribed-backend/internal/app/command"
	"github.com/subscribeddotdev/subscribed-backend/internal/common/messaging"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (p Publisher) PublishMessageSent(ctx context.Context, e command.MessageSent) error {
	evt := events.MessageSent{
		Header: &events.Header{
			Id:            domain.NewID().String(),
			Name:          command.MessageSentEvent,
			CorrelationId: "", // TODO: retrieve it from ctx
			PublisherName: "subscribed-backend",
			PublishedAt:   timestamppb.New(time.Now().UTC()),
		},
		MessageId:  e.MessageID,
		EndpointId: e.EndpointID,
	}

	payload, err := json.Marshal(&evt)
	if err != nil {
		return fmt.Errorf("error to marshall event %s due to: %v", command.MessageSentEvent, err)
	}

	err = p.eventPublisher.Publish(command.MessageSentEvent, message.Message{
		Payload: payload,
	})
	if err != nil {
		return fmt.Errorf("error publishing event %s due to: %v", command.MessageSentEvent, err)
	}

	return nil
}
