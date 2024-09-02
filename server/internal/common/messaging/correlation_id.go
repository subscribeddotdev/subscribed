package messaging

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/subscribeddotdev/subscribed-backend/internal/common/observability"
)

const correlationIdContextKey = "correlation_id"

type CorrelationIdPublisher struct {
	publisher message.Publisher
}

func (c CorrelationIdPublisher) Publish(topic string, messages ...*message.Message) error {
	for _, msg := range messages {
		correlationID, err := observability.CorrelationIdFromContext(msg.Context())
		if err == nil {
			msg.Metadata.Set(correlationIdContextKey, correlationID)
		}
	}

	return c.publisher.Publish(topic, messages...)
}

func (c CorrelationIdPublisher) Close() error {
	return c.publisher.Close()
}

func NewCorrelationIdPublisher(publisher message.Publisher) CorrelationIdPublisher {
	return CorrelationIdPublisher{
		publisher: publisher,
	}
}
