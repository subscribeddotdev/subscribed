package iam

import (
	"time"

	"github.com/friendsofgo/errors"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type Organization struct {
	id         domain.ID
	createdAt  time.Time
	disabledAt *time.Time
	users      []Member
}

func NewOrganization() *Organization {
	return &Organization{
		id:         domain.NewID(),
		createdAt:  time.Now(),
		disabledAt: nil,
	}
}

func MarshallToOrganization(
	id domain.ID,
	createdAt time.Time,
	disabledAt *time.Time,
) (*Organization, error) {
	if id.IsEmpty() {
		return nil, errors.New("id cannot be empty")
	}

	if createdAt.After(time.Now()) {
		return nil, errors.New("createdAt cannot be set in the future")
	}

	if disabledAt != nil && disabledAt.After(time.Now()) {
		return nil, errors.New("disabledAt cannot be set in the future")
	}

	return &Organization{
		id:         id,
		createdAt:  createdAt,
		disabledAt: disabledAt,
		users:      nil,
	}, nil
}

func (o *Organization) Id() domain.ID {
	return o.id
}
