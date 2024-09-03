package messaging

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/subscribeddotdev/subscribed/server/internal/common/logs"
	"github.com/subscribeddotdev/subscribed/server/internal/common/observability"
)

func CorrelationIdMiddleware(h message.HandlerFunc) message.HandlerFunc {
	return func(msg *message.Message) ([]*message.Message, error) {
		if msg.Metadata.Get("correlation_id") != "" {
			msgCtx := observability.ContextWithCorrelationID(msg.Context(), msg.Metadata.Get("correlation_id"))
			msg.SetContext(msgCtx)
		}

		return h(msg)
	}
}

func ErrorLoggerMiddleware(logger *logs.Logger) func(h message.HandlerFunc) message.HandlerFunc {
	return func(h message.HandlerFunc) message.HandlerFunc {
		return func(msg *message.Message) ([]*message.Message, error) {
			logger.Info(string(msg.Payload))

			msgs, err := h(msg)
			if err != nil {
				logger.Error("an error occurred while processing message", "payload", string(msg.Payload), "error", err)
			}

			return msgs, err
		}
	}
}
