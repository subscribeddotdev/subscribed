package iam

import (
	"errors"
	"strings"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type LoginProviderID string

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
	}, nil
}
