package domain

import (
	"errors"
	"strings"
	"time"
)

type ApplicationID string

func (i ApplicationID) String() string {
	return string(i)
}

func NewApplicationID() ApplicationID {
	return ApplicationID(NewID().WithPrefix("app"))
}

type Application struct {
	id        ApplicationID
	name      string
	envID     EnvironmentID
	createdAt time.Time
}

func NewApplication(name string, envID EnvironmentID) (*Application, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if envID.String() == "" {
		return nil, errors.New("envID cannot be empty")
	}

	return &Application{
		id:        NewApplicationID(),
		name:      name,
		envID:     envID,
		createdAt: time.Now(),
	}, nil
}

func (a *Application) EnvID() EnvironmentID {
	return a.envID
}

func (a *Application) ID() ApplicationID {
	return a.id
}

func (a *Application) Name() string {
	return a.name
}

func (a *Application) CreatedAt() time.Time {
	return a.createdAt
}

func UnMarshallApplication(id ApplicationID, name string, envID EnvironmentID, createdAt time.Time) *Application {
	return &Application{
		id:        id,
		name:      name,
		envID:     envID,
		createdAt: createdAt,
	}
}
