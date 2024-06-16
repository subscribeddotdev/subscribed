package domain

import (
	"time"
)

type Environment struct {
	id           ID
	name         string
	orgID        ID
	createdAt    time.Time
	archivedAt   *time.Time
	apiKeys      []ApiKey
	applications []Application
	eventTypes   []EventType
}

func (e *Environment) Id() ID {
	return e.id
}

func (e *Environment) Name() string {
	return e.name
}

func (e *Environment) OrgID() ID {
	return e.orgID
}

func (e *Environment) CreatedAt() time.Time {
	return e.createdAt
}

func (e *Environment) ArchivedAt() *time.Time {
	return e.archivedAt
}

func (e *Environment) ApiKeys() []ApiKey {
	return e.apiKeys
}

func (e *Environment) Applications() []Application {
	return e.applications
}

func (e *Environment) EventTypes() []EventType {
	return e.eventTypes
}
