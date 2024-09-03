package command

const (
	MessageSentEvent = "MessageSent"
)

type MessageSent struct {
	MessageID  string
	EndpointID string
}
