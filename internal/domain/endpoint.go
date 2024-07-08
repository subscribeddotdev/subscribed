package domain

import (
	"time"

	"github.com/friendsofgo/errors"
)

type Headers map[string]string

type EndpointID string

func (i EndpointID) String() string {
	return string(i)
}

func NewEndpointID() EndpointID {
	return EndpointID(NewID().WithPrefix("end"))
}

type Endpoint struct {
	id            EndpointID
	url           EndpointURL
	applicationID ApplicationID
	description   string
	headers       Headers
	eventTypeIDs  []EventTypeID
	signingSecret SigningSecret
	createdAt     time.Time
	updatedAt     time.Time
}

func NewEndpoint(
	endpointURL EndpointURL,
	applicationID ApplicationID,
	description string,
	eventTypeIDs []EventTypeID,
) (*Endpoint, error) {
	if endpointURL.IsEmpty() {
		return nil, errors.New("endpointURL cannot be empty")
	}

	if applicationID.String() == "" {
		return nil, errors.New("applicationID cannot be empty")
	}

	signingSecret, err := NewSigningSecret()
	if err != nil {
		return nil, err
	}

	return &Endpoint{
		id:            NewEndpointID(),
		url:           endpointURL,
		applicationID: applicationID,
		description:   description,
		eventTypeIDs:  eventTypeIDs,
		createdAt:     time.Now().UTC(),
		updatedAt:     time.Now().UTC(),
		signingSecret: signingSecret,
		headers:       nil,
	}, nil
}

func (e *Endpoint) ID() EndpointID {
	return e.id
}

func (e *Endpoint) EndpointURL() EndpointURL {
	return e.url
}

func (e *Endpoint) Description() string {
	return e.description
}

func (e *Endpoint) EventTypeIDs() []EventTypeID {
	return e.eventTypeIDs
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

func (e *Endpoint) ApplicationID() ApplicationID {
	return e.applicationID
}

func UnMarshallEndpoint(
	id EndpointID,
	applicationID ApplicationID,
	endpointURL,
	description string,
	headers Headers,
	eventTypeIDs []EventTypeID,
	signingSecret string,
	createdAt,
	updatedAt time.Time,
) (*Endpoint, error) {
	dEndpointURL, err := NewEndpointURL(endpointURL)
	if err != nil {
		return nil, err
	}

	ss, err := UnMarshallSigningSecret(signingSecret)
	if err != nil {
		return nil, err
	}

	return &Endpoint{
		id:            id,
		url:           dEndpointURL,
		applicationID: applicationID,
		description:   description,
		headers:       headers,
		eventTypeIDs:  eventTypeIDs,
		signingSecret: ss,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
	}, nil
}
