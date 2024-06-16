package domain

import "time"

type ApiKey struct {
	id        ID
	name      string
	key       string
	createAt  time.Time
	expiresAt *time.Time
}

func (a *ApiKey) Id() ID {
	return a.id
}

func (a *ApiKey) Name() string {
	return a.name
}

func (a *ApiKey) Key() string {
	return a.key
}

func (a *ApiKey) CreateAt() time.Time {
	return a.createAt
}

func (a *ApiKey) ExpiresAt() *time.Time {
	return a.expiresAt
}
