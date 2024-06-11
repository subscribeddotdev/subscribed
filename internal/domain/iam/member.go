package iam

import (
	"errors"
	"strings"
	"time"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type LoginProviderID string

func (i LoginProviderID) String() string {
	return string(i)
}

func (i LoginProviderID) Validate() error {
	if strings.TrimSpace(string(i)) == "" {
		return errors.New("loginProviderID cannot be empty")
	}

	return nil
}

type Member struct {
	id              domain.ID
	organizationID  domain.ID
	loginProviderId LoginProviderID
	firstName       string
	lastName        string
	email           Email
	createdAt       time.Time
}

func NewMember(
	organizationID domain.ID,
	loginProviderId LoginProviderID,
	firstName,
	lastName string,
	email Email,
) (*Member, error) {
	if organizationID.IsEmpty() {
		return nil, errors.New("organizationID cannot be empty")
	}

	if email.IsEmpty() {
		return nil, errors.New("email cannot be empty")
	}

	if err := loginProviderId.Validate(); err != nil {
		return nil, err
	}

	return &Member{
		id:              domain.NewID(),
		organizationID:  organizationID,
		loginProviderId: loginProviderId,
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		createdAt:       time.Now().UTC(),
	}, nil
}

func (m *Member) Id() domain.ID {
	return m.id
}

func (m *Member) OrganizationID() domain.ID {
	return m.organizationID
}

func (m *Member) LoginProviderId() LoginProviderID {
	return m.loginProviderId
}

func (m *Member) FirstName() string {
	return m.firstName
}

func (m *Member) LastName() string {
	return m.lastName
}

func (m *Member) Email() Email {
	return m.email
}

func (m *Member) CreatedAt() time.Time {
	return m.createdAt
}
