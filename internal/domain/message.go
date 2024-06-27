package domain

import (
	"errors"
	"strings"
	"time"
)

type Message struct {
	id           ID
	eventTypeID  ID
	sentAt       time.Time
	payload      string
	sendAttempts []MessageSendAttempt
}

func NewMessage(eventTypeID ID, payload string) (*Message, error) {
	if eventTypeID.IsEmpty() {
		return nil, errors.New("eventTypeID cannot be empty")
	}

	if strings.TrimSpace(payload) == "" {
		return nil, errors.New("payload cannot be empty")
	}

	return &Message{
		id:           NewID(),
		eventTypeID:  eventTypeID,
		sentAt:       time.Now().UTC(),
		payload:      payload,
		sendAttempts: nil,
	}, nil
}

func (m *Message) Id() ID {
	return m.id
}

func (m *Message) EventTypeID() ID {
	return m.eventTypeID
}

func (m *Message) CreatedAt() time.Time {
	return m.sentAt
}

func (m *Message) Payload() string {
	return m.payload
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
