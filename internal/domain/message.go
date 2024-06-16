package domain

import "time"

type Message struct {
	id           ID
	eventType    EventType
	createdAt    time.Time
	content      string
	sendAttempts []MessageSendAttempt
}

func (m *Message) Id() ID {
	return m.id
}

func (m *Message) EventType() EventType {
	return m.eventType
}

func (m *Message) CreatedAt() time.Time {
	return m.createdAt
}

func (m *Message) Content() string {
	return m.content
}

func (m *Message) SendAttempts() []MessageSendAttempt {
	return m.sendAttempts
}

type MessageSendAttempt struct {
	timestamp  time.Time
	status     string
	response   string
	statusCode int
	headers    map[string]string
}

func (m *MessageSendAttempt) Timestamp() time.Time {
	return m.timestamp
}

func (m *MessageSendAttempt) Status() string {
	return m.status
}

func (m *MessageSendAttempt) Response() string {
	return m.response
}

func (m *MessageSendAttempt) StatusCode() int {
	return m.statusCode
}

func (m *MessageSendAttempt) Headers() map[string]string {
	return m.headers
}
