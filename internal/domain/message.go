package domain

import (
	"errors"
	"strings"
	"time"
)

type MessageID string

func NewMessageID() MessageID {
	return MessageID(NewID().WithPrefix("msg").String())
}

func (i MessageID) String() string {
	return string(i)
}

type Message struct {
	id            MessageID
	eventTypeID   EventTypeID
	applicationID ApplicationID
	orgID         string
	sentAt        time.Time
	payload       string
	sendAttempts  []MessageSendAttempt
}

func NewMessage(eventTypeID EventTypeID, orgID string, applicationID ApplicationID, payload string) (*Message, error) {
	if eventTypeID.String() == "" {
		return nil, errors.New("eventTypeID cannot be empty")
	}

	if applicationID.String() == "" {
		return nil, errors.New("applicationID cannot be empty")
	}

	if orgID == "" {
		return nil, errors.New("orgID cannot be empty")
	}

	if strings.TrimSpace(payload) == "" {
		return nil, errors.New("payload cannot be empty")
	}

	return &Message{
		id:            NewMessageID(),
		orgID:         orgID,
		eventTypeID:   eventTypeID,
		applicationID: applicationID,
		sentAt:        time.Now().UTC(),
		payload:       payload,
		sendAttempts:  nil,
	}, nil
}

func (m *Message) Id() MessageID {
	return m.id
}

func (m *Message) EventTypeID() EventTypeID {
	return m.eventTypeID
}

func (m *Message) OrgID() string {
	return m.orgID
}

func (m *Message) ApplicationID() ApplicationID {
	return m.applicationID
}

func (m *Message) SentAt() time.Time {
	return m.sentAt
}

func (m *Message) Payload() string {
	return m.payload
}

func (m *Message) SendAttempts() []MessageSendAttempt {
	return m.sendAttempts
}

func UnMarshallMessage(
	id MessageID,
	eventTypeID EventTypeID,
	applicationID ApplicationID,
	orgID string,
	sentAt time.Time,
	payload string,
) (*Message, error) {
	return &Message{
		id:            id,
		eventTypeID:   eventTypeID,
		applicationID: applicationID,
		orgID:         orgID,
		sentAt:        sentAt,
		payload:       payload,
	}, nil
}
