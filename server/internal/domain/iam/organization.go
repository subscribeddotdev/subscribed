package iam

import (
	"time"

	"github.com/friendsofgo/errors"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type OrgID string

func (i OrgID) String() string {
	return string(i)
}

func NewOrgID() OrgID {
	return OrgID(domain.NewID().WithPrefix("org"))
}

type Organization struct {
	id         OrgID
	createdAt  time.Time
	disabledAt *time.Time
	users      []Member
}

func NewOrganization() *Organization {
	return &Organization{
		id:         NewOrgID(),
		createdAt:  time.Now(),
		disabledAt: nil,
	}
}

func MarshallToOrganization(
	id OrgID,
	createdAt time.Time,
	disabledAt *time.Time,
) (*Organization, error) {
	if id.String() == "" {
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

func (o *Organization) ID() OrgID {
	return o.id
}

func (o *Organization) CreatedAt() time.Time {
	return o.createdAt
}

func (o *Organization) DisabledAt() *time.Time {
	return o.disabledAt
}
