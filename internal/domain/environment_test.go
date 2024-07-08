package domain_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

func TestNewEnvironment(t *testing.T) {
	testCases := []struct {
		name        string
		expectedErr string

		envName string
		envType domain.EnvType
		orgID   string
	}{
		{
			name:        "create_new_environment",
			expectedErr: "",
			envName:     "Development",
			orgID:       iam.NewOrgID().String(),
			envType:     domain.EnvTypeDevelopment,
		},
		{
			name:        "error_empty_name",
			expectedErr: "name cannot be empty",
			envName:     "  ",
			orgID:       iam.NewOrgID().String(),
			envType:     domain.EnvTypeDevelopment,
		},
		{
			name:        "error_empty_or_invalid_org_id",
			expectedErr: "orgID cannot be empty",
			envName:     "Production",
			orgID:       "",
			envType:     domain.EnvTypeProduction,
		},
		{
			name:        "error_empty_env_type",
			expectedErr: "environment type cannot be empty",
			envName:     "Production",
			orgID:       iam.NewOrgID().String(),
			envType:     domain.EnvType{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			env, err := domain.NewEnvironment(tc.envName, tc.orgID, tc.envType)

			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.envName, env.Name())
			assert.Equal(t, tc.orgID, env.OrgID())
			assert.Equal(t, tc.envType, env.Type())
			assert.True(t, env.CreatedAt().Before(time.Now()))
		})
	}
}
