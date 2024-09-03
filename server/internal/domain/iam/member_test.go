package iam_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed/server/internal/domain/iam"
	"github.com/subscribeddotdev/subscribed/server/tests"
)

func TestNewMember(t *testing.T) {
	testCases := []struct {
		name        string
		expectedErr string

		organizationID iam.OrgID
		firstName      string
		lastName       string
		email          iam.Email
		password       iam.Password
	}{
		{
			name:           "new_member",
			expectedErr:    "",
			organizationID: iam.NewOrgID(),
			firstName:      gofakeit.FirstName(),
			lastName:       gofakeit.LastName(),
			email:          tests.MustEmail(t, gofakeit.Email()),
			password:       tests.FixturePassword(t),
		},
		{
			name:           "error_empty_organization_id",
			expectedErr:    "organizationID cannot be empty",
			organizationID: iam.OrgID(""),
			firstName:      gofakeit.FirstName(),
			lastName:       gofakeit.LastName(),
			email:          tests.MustEmail(t, gofakeit.Email()),
			password:       tests.FixturePassword(t),
		},
		{
			name:           "error_empty_email_address",
			expectedErr:    "email cannot be empty",
			organizationID: iam.NewOrgID(),
			firstName:      gofakeit.FirstName(),
			lastName:       gofakeit.LastName(),
			email:          iam.Email{},
			password:       tests.FixturePassword(t),
		},
		{
			name:           "error_empty_password",
			expectedErr:    "password cannot be empty",
			organizationID: iam.NewOrgID(),
			firstName:      gofakeit.FirstName(),
			lastName:       gofakeit.LastName(),
			email:          tests.MustEmail(t, gofakeit.Email()),
			password:       iam.Password{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			member, err := iam.NewMember(
				tc.organizationID,
				tc.firstName,
				tc.lastName,
				tc.email,
				tc.password,
			)

			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.firstName, member.FirstName())
			assert.Equal(t, tc.lastName, member.LastName())
			assert.Equal(t, tc.email, member.Email())
			assert.Equal(t, tc.password.String(), member.Password().String())
		})
	}
}
