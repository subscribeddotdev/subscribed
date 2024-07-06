package domain_test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

func TestNewApplication(t *testing.T) {
	testCases := []struct {
		name        string
		expectedErr string

		appName string
		envID   domain.ID
	}{
		{
			name:        "create_new_application",
			expectedErr: "",
			appName:     gofakeit.AppName(),
			envID:       domain.NewID(),
		},
		{
			name:        "error_empty_name",
			expectedErr: "name cannot be empty",
			appName:     " ",
			envID:       domain.NewID(),
		},
		{
			name:        "error_invalid_or_empty_env_id",
			expectedErr: "envID cannot be empty",
			appName:     gofakeit.AppName(),
			envID:       domain.ID(""),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			application, err := domain.NewApplication(tc.appName, tc.envID)

			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				return
			}

			require.NoError(t, err)

			assert.NotNil(t, application.Id())
			assert.Equal(t, tc.appName, application.Name())
			assert.Equal(t, tc.envID, application.EnvID())
			assert.True(t, application.CreatedAt().Before(time.Now()))
		})
	}
}
