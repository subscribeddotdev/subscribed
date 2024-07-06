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
	eventTypeID   ID
	applicationID ID
	orgID         ID
	sentAt        time.Time
	payload       string
	sendAttempts  []MessageSendAttempt
}

func NewMessage(eventTypeID, orgID, applicationID ID, payload string) (*Message, error) {
	if eventTypeID.IsEmpty() {
		return nil, errors.New("eventTypeID cannot be empty")
	}

	if applicationID.IsEmpty() {
		return nil, errors.New("applicationID cannot be empty")
	}

	if orgID.IsEmpty() {
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

func (m *Message) EventTypeID() ID {
	return m.eventTypeID
}

func (m *Message) OrgID() ID {
	return m.orgID
}

func (m *Message) ApplicationID() ID {
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
	id,
	eventTypeID,
	applicationID,
	orgID string,
	sentAt time.Time,
	payload string,
) (*Message, error) {
	dID, err := NewIdFromString(id)
	if err != nil {
		return nil, err
	}

	dEventTypeID, err := NewIdFromString(eventTypeID)
	if err != nil {
		return nil, err
	}

	dApplicationID, err := NewIdFromString(applicationID)
	if err != nil {
		return nil, err
	}

	dOrgID, err := NewIdFromString(orgID)
	if err != nil {
		return nil, err
	}

	return &Message{
		id:            MessageID(dID),
		eventTypeID:   dEventTypeID,
		applicationID: dApplicationID,
		orgID:         dOrgID,
		sentAt:        sentAt,
		payload:       payload,
	}, nil
}
