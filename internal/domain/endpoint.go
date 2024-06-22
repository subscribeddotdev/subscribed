package domain

import (
	"fmt"
	"time"

	"github.com/friendsofgo/errors"
)

// SigningSecret TODO: pending implementation
type SigningSecret string

func (s SigningSecret) String() string {
	return string(s)
}

type Headers map[string]string

type Endpoint struct {
	id                     ID
	url                    EndpointURL
	applicationID          ID
	description            string
	headers                Headers
	eventTypesSubscribedTo []EventType
	signingSecret          SigningSecret
	createdAt              time.Time
	updatedAt              time.Time
}

func NewEndpoint(
	endpointURL EndpointURL,
	applicationID ID,
	description string,
	eventTypesSubscribedTo []EventType,
) (*Endpoint, error) {
	if endpointURL.IsEmpty() {
		return nil, errors.New("endpointURL cannot be empty")
	}

	if applicationID.IsEmpty() {
		return nil, errors.New("applicationID cannot be empty")
	}

	return &Endpoint{
		id:                     NewID(),
		url:                    endpointURL,
		applicationID:          applicationID,
		description:            description,
		eventTypesSubscribedTo: eventTypesSubscribedTo,
		createdAt:              time.Now().UTC(),
		updatedAt:              time.Now().UTC(),
		signingSecret:          SigningSecret(fmt.Sprintf("whsec_%s", NewID())), //TODO: replace this with a hashed secret
		headers:                nil,
	}, nil
}

func (e *Endpoint) Id() ID {
	return e.id
}

func (e *Endpoint) EndpointURL() EndpointURL {
	return e.url
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

func (e *Endpoint) SigningSecret() SigningSecret {
	return e.signingSecret
}

func (e *Endpoint) Headers() map[string]string {
	return e.headers
}

func (e *Endpoint) ApplicationID() ID {
	return e.applicationID
}
