package domain

import (
	"errors"
	"strings"
	"time"
)

type Message struct {
	id            ID
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
		id:            NewID(),
		orgID:         orgID,
		eventTypeID:   eventTypeID,
		applicationID: applicationID,
		sentAt:        time.Now().UTC(),
		payload:       payload,
		sendAttempts:  nil,
	}, nil
}

func (m *Message) Id() ID {
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
		id:            dID,
		eventTypeID:   dEventTypeID,
		applicationID: dApplicationID,
		orgID:         dOrgID,
		sentAt:        sentAt,
		payload:       payload,
	}, nil
}

type MessageSendAttempt struct {
	id         ID
	endpointID ID
	timestamp  time.Time
	status     string
	response   string
	statusCode uint
	headers    Headers
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

func (m *MessageSendAttempt) StatusCode() uint {
	return m.statusCode
}

func (m *MessageSendAttempt) Headers() Headers {
	return m.headers
}
