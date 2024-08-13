package iam_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

func TestNewPassword(t *testing.T) {
	testCases := []struct {
		name        string
		expectedErr string

		plainText string
	}{
		{
			name:        "create_new_password",
			expectedErr: "",
			plainText:   gofakeit.Password(true, true, true, true, false, 12),
		},
		{
			name:        "error_empty_password",
			expectedErr: "password cannot be empty",
			plainText:   "                    ",
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			password, err := iam.NewPassword(tc.plainText)
			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				return
			}

			require.NoError(t, err)
			assert.NotEmpty(t, password.String())
		})
	}
}

func TestNewPasswordFromHash(t *testing.T) {
	testCases := []struct {
		name        string
		expectedErr string

		hash string
	}{
		{
			name:        "create_new_password_from_hash",
			expectedErr: "",
			hash:        fixtureHash(t),
		},
		{
			name:        "error_invalid_hash",
			expectedErr: "the hash provided is invalid",
			hash:        gofakeit.Sentence(10),
		},
		{
			name:        "error_invalid_hash_empty_string",
			expectedErr: "the hash provided is invalid",
			hash:        "         ",
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			password, err := iam.NewPasswordFromHash(tc.hash)
			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				return
			}

			require.NoError(t, err)
			assert.NotEmpty(t, password.String())
		})
	}
}

func fixtureHash(t *testing.T) string {
	p, err := iam.NewPassword(gofakeit.Password(true, true, true, true, false, 12))
	require.NoError(t, err)
	return p.String()
}
