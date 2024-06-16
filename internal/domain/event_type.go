package domain

import "time"

type EventType struct {
	id          ID
	name        string
	description string
	//TODO: JSON http://json-schema.org
	schema        string
	schemaExample string
	archivedAt    *time.Time
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
