package iam_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

func TestNewEmail(t *testing.T) {
	testCases := []struct {
		name        string
		expectedErr string

		organizationID  domain.ID
		loginProviderId iam.LoginProviderID
		firstName       string
		lastName        string
		email           iam.Email
	}{
		{
			name:            "new_member",
			expectedErr:     "",
			organizationID:  domain.NewID(),
			loginProviderId: iam.LoginProviderID(gofakeit.UUID()),
			firstName:       gofakeit.FirstName(),
			lastName:        gofakeit.LastName(),
			email:           mustEmail(t, gofakeit.Email()),
		},
		{
			name:            "error_empty_organization_id",
			expectedErr:     "organizationID cannot be empty",
			organizationID:  domain.ID{},
			loginProviderId: iam.LoginProviderID(gofakeit.UUID()),
			firstName:       gofakeit.FirstName(),
			lastName:        gofakeit.LastName(),
			email:           mustEmail(t, gofakeit.Email()),
		},
		{
			name:            "error_empty_email_address",
			expectedErr:     "email cannot be empty",
			organizationID:  domain.NewID(),
			loginProviderId: iam.LoginProviderID(gofakeit.UUID()),
			firstName:       gofakeit.FirstName(),
			lastName:        gofakeit.LastName(),
			email:           iam.Email{},
		},
		{
			name:            "error_empty_login_provider_id",
			expectedErr:     "loginProviderID cannot be empty",
			organizationID:  domain.NewID(),
			loginProviderId: iam.LoginProviderID(""),
			firstName:       gofakeit.FirstName(),
			lastName:        gofakeit.LastName(),
			email:           mustEmail(t, gofakeit.Email()),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			member, err := iam.NewMember(
				tc.organizationID,
				tc.loginProviderId,
				tc.firstName,
				tc.lastName,
				tc.email,
			)

			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.loginProviderId, member.LoginProviderId())
			assert.Equal(t, tc.firstName, member.FirstName())
			assert.Equal(t, tc.lastName, member.LastName())
			assert.Equal(t, tc.email, member.Email())
		})
	}
}

func mustEmail(t *testing.T, address string) iam.Email {
	email, err := iam.NewEmail(gofakeit.Email())
	require.NoError(t, err)
	return email
}
