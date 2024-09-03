package iam_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed/server/internal/domain/iam"
)

func TestNewEmail(t *testing.T) {
	testCases := []struct {
		name        string
		expectedErr string

		address string
	}{
		{
			name:        "valid_email_address",
			expectedErr: "",
			address:     "john.doe@gmail.com",
		},
		{
			name:        "error_missing_@_sign",
			expectedErr: "mail: missing '@' or angle-addr",
			address:     "john.doegmail.com",
		},
		{
			name:        "error_empty_address",
			expectedErr: "mail: no address",
			address:     " ",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			email, err := iam.NewEmail(tc.address)

			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.address, email.String())
			assert.False(t, email.IsEmpty())
		})
	}
}
