package domain

import (
	"errors"
	"time"
)

const httpLastSuccessStatusCode = 299

type StatusCode uint

type SendAttemptStatus string

func (s SendAttemptStatus) String() string {
	return string(s)
}

/*func (s SendAttemptStatus) validate() error {
	if s != "succeeded" && s != "failed" {
		return fmt.Errorf("%s is not a valid message send attempt status", s)
	}

	return nil
}*/

var (
	sendAttemptStatusSucceeded = SendAttemptStatus("succeeded")
	sendAttemptStatusFailed    = SendAttemptStatus("failed")
)

type MessageSendAttempt struct {
	id         ID
	messageID  ID
	endpointID ID
	timestamp  time.Time
	status     SendAttemptStatus
	response   string // TODO: make this value-object
	statusCode StatusCode
	headers    Headers
}

func NewMessageSendAttempt(
	endpointID, messageID ID,
	response string,
	statusCode StatusCode,
	headers Headers,
) (*MessageSendAttempt, error) {
	if endpointID.IsEmpty() {
		return nil, errors.New("endpointURL cannot be empty")
	}

	if messageID.IsEmpty() {
		return nil, errors.New("messageID cannot be empty")
	}

	status := sendAttemptStatusFailed
	if statusCode <= httpLastSuccessStatusCode {
		status = sendAttemptStatusSucceeded
	}

	return &MessageSendAttempt{
		id:         NewID(),
		messageID:  messageID,
		endpointID: endpointID,
		timestamp:  time.Now().UTC(),
		status:     status,
		response:   response,
		statusCode: statusCode,
		headers:    headers,
	}, nil
}

func (m *MessageSendAttempt) Id() ID {
	return m.id
}

func (m *MessageSendAttempt) MessageID() ID {
	return m.messageID
}

func (m *MessageSendAttempt) EndpointID() ID {
	return m.endpointID
}

func (m *MessageSendAttempt) Timestamp() time.Time {
	return m.timestamp
}

func (m *MessageSendAttempt) Status() SendAttemptStatus {
	return m.status
}

func (m *MessageSendAttempt) Response() string {
	return m.response
}

func (m *MessageSendAttempt) StatusCode() StatusCode {
	return m.statusCode
}

func (m *MessageSendAttempt) Headers() Headers {
	return m.headers
}
