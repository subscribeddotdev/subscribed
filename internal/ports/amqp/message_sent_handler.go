package amqp

import (
	"encoding/json"
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
	events "github.com/subscribeddotdev/subscribed-backend/events/go"
	"github.com/subscribeddotdev/subscribed-backend/internal/app"
	"github.com/subscribeddotdev/subscribed-backend/internal/app/command"
)

type MessageSentHandler struct {
	application *app.App
}

func (h MessageSentHandler) HandlerName() string {
	return command.MessageSentEvent + "_Handler"
}

func (h MessageSentHandler) EventName() string {
	return command.MessageSentEvent
}

func (h MessageSentHandler) Handle(m *message.Message) error {
	var payload events.MessageSent
	err := json.Unmarshal(m.Payload, &payload)
	if err != nil {
		return fmt.Errorf("error unmarshalling event '%s': %v", command.MessageSentEvent, err)
	}

	fmt.Println(payload.MessageId, payload.EndpointId, m.Context())

	return h.application.Command.CallWebhookEndpoint.Execute(m.Context(), command.CallWebhookEndpoint{
		EndpointID: payload.EndpointId,
		MessageID:  payload.MessageId,
	})
}
