package domain

import (
	"strings"
	"time"

	"errors"
)

type EventTypeID string

func (i EventTypeID) String() string {
	return string(i)
}

func NewEventTypeID() EventTypeID {
	return EventTypeID(NewID().WithPrefix("ety"))
}

type EventType struct {
	id          EventTypeID
	orgID       string
	name        string
	description string
	//TODO: JSON http://json-schema.org
	schema        string
	schemaExample string
	createdAt     time.Time
	archivedAt    *time.Time
}

func NewEventType(orgID, name, description, schema, schemaExample string) (*EventType, error) {
	orgID = strings.TrimSpace(orgID)
	if orgID == "" {
		return nil, errors.New("orgID cannot be empty")
	}

	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	return &EventType{
		id:            NewEventTypeID(),
		orgID:         orgID,
		name:          name,
		description:   description,
		schema:        schema,
		schemaExample: schemaExample,
		archivedAt:    nil,
		createdAt:     time.Now().UTC(),
	}, nil
}

func (e *EventType) OrgID() string {
	return e.orgID
}

func (e *EventType) ID() EventTypeID {
	return e.id
}

func (e *EventType) Name() string {
	return e.name
}

func (e *EventType) Description() string {
	return e.description
}

func (e *EventType) Schema() string {
	return e.schema
}

func (e *EventType) SchemaExample() string {
	return e.schemaExample
}

func (e *EventType) ArchivedAt() *time.Time {
	return e.archivedAt
}

func (e *EventType) CreatedAt() time.Time {
	return e.createdAt
}

func UnMarshallEventType(
	id EventTypeID,
	orgID string,
	name, description, schema, schemaExample string,
	createdAt time.Time,
	archivedAt *time.Time,
) *EventType {
	return &EventType{
		id:            id,
		orgID:         orgID,
		name:          name,
		description:   description,
		schema:        schema,
		schemaExample: schemaExample,
		createdAt:     createdAt.UTC(),
		archivedAt:    archivedAt,
	}
}
