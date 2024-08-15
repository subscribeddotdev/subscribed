package iam

import (
	"errors"
	"fmt"
	"time"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrAuthenticationFailed = errors.New("error authenticating member")
	ErrMemberAlreadyExists  = errors.New("error member already exists")
)

type MemberID string

func (i MemberID) String() string {
	return string(i)
}

func NewMemberID() MemberID {
	return MemberID(domain.NewID().WithPrefix("mem"))
}

type Member struct {
	id             MemberID
	organizationID OrgID
	firstName      string
	lastName       string
	email          Email
	password       Password
	createdAt      time.Time
}

func NewMember(
	organizationID OrgID,
	firstName,
	lastName string,
	email Email,
	password Password,
) (*Member, error) {
	if organizationID.String() == "" {
		return nil, errors.New("organizationID cannot be empty")
	}

	if email.IsEmpty() {
		return nil, errors.New("email cannot be empty")
	}

	if password.IsEmpty() {
		return nil, errors.New("password cannot be empty")
	}

	return &Member{
		id:             NewMemberID(),
		organizationID: organizationID,
		firstName:      firstName,
		lastName:       lastName,
		email:          email,
		password:       password,
		createdAt:      time.Now(),
	}, nil
}

func (m *Member) ID() MemberID {
	return m.id
}

func (m *Member) OrgID() OrgID {
	return m.organizationID
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

func (m *Member) Password() Password {
	return m.password
}

func (m *Member) Authenticate(plainTextPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(m.password.hash), []byte(plainTextPassword))
	if err != nil {
		return ErrAuthenticationFailed
	}

	return nil
}

func (m *Member) CreatedAt() time.Time {
	return m.createdAt
}

func UnMarshallMember(
	id MemberID,
	orgID OrgID,
	firstName, lastName, email, password string,
	createdAt time.Time,
) (*Member, error) {
	mEmail, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	if time.Now().Before(createdAt) {
		return nil, fmt.Errorf("createdAt '%s' is set in the future", createdAt)
	}

	dPassword, err := NewPasswordFromHash(password)
	if err != nil {
		return nil, err
	}

	return &Member{
		id:             id,
		organizationID: orgID,
		firstName:      firstName,
		lastName:       lastName,
		email:          mEmail,
		password:       dPassword,
		createdAt:      createdAt,
	}, nil
}
