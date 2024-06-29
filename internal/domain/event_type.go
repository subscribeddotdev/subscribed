package domain

import (
	"strings"
	"time"

	"errors"
)

type EventType struct {
	id          ID
	orgID       ID
	name        string
	description string
	//TODO: JSON http://json-schema.org
	schema        string
	schemaExample string
	createdAt     time.Time
	archivedAt    *time.Time
}

func NewEventType(orgID ID, name, description, schema, schemaExample string) (*EventType, error) {
	if orgID.IsEmpty() {
		return nil, errors.New("orgID cannot be empty")
	}

	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	return &EventType{
		id:            NewID(),
		orgID:         orgID,
		name:          name,
		description:   description,
		schema:        schema,
		schemaExample: schemaExample,
		archivedAt:    nil,
		createdAt:     time.Now().UTC(),
	}, nil
}

func (e *EventType) OrgID() ID {
	return e.orgID
}

func (e *EventType) Id() ID {
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
