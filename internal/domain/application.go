package domain

import (
	"strings"
	"time"

	"github.com/friendsofgo/errors"
)

type Application struct {
	id        ID
	name      string
	envID     ID
	createdAt time.Time
	endpoints []Endpoint
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
		endpoints: nil,
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

func (a *Application) Endpoints() []Endpoint {
	return a.endpoints
}
