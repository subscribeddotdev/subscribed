package domain

import (
	"net/url"
	"time"
)

type Endpoint struct {
	id                     ID
	endpointUrl            url.URL
	description            string
	eventTypesSubscribedTo []EventType
	createdAt              time.Time
	updatedAt              time.Time
	signingSecret          string
	headers                map[string]string
	messages               []Message
}

func (e *Endpoint) Id() ID {
	return e.id
}

func (e *Endpoint) EndpointUrl() url.URL {
	return e.endpointUrl
}

func (e *Endpoint) Description() string {
	return e.description
}

func (e *Endpoint) EventTypesSubscribedTo() []EventType {
	return e.eventTypesSubscribedTo
}

func (e *Endpoint) CreatedAt() time.Time {
	return e.createdAt
}

func (e *Endpoint) UpdatedAt() time.Time {
	return e.updatedAt
}

func (e *Endpoint) SigningSecret() string {
	return e.signingSecret
}

func (e *Endpoint) Headers() map[string]string {
	return e.headers
}

func (e *Endpoint) Messages() []Message {
	return e.messages
}
