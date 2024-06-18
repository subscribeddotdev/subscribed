package domain

import (
	"errors"
	"strings"
	"time"
)

type Application struct {
	id        ID
	name      string
	envID     ID
	createdAt time.Time
}

func NewApplication(name string, envID ID) (*Application, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if envID.IsEmpty() {
		return nil, errors.New("envID cannot be empty")
	}

	return &Application{
		id:        NewID(),
		name:      name,
		envID:     envID,
		createdAt: time.Now().UTC(),
	}, nil
}

func (a *Application) EnvID() ID {
	return a.envID
}

func (a *Application) Id() ID {
	return a.id
}

func (a *Application) Name() string {
	return a.name
}

func (a *Application) CreatedAt() time.Time {
	return a.createdAt
}
