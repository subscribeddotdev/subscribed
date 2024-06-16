package domain

import "time"

type Application struct {
	id        ID
	name      string
	createdAt time.Time
	endpoints []Endpoint
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
