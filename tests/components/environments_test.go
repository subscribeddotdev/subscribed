package components_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/tests"
	"github.com/subscribeddotdev/subscribed-backend/tests/fixture"
	"github.com/subscribeddotdev/subscribed-backend/tests/jwks"
)

func TestEnvironments(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	member := ff.NewMember().WithOrganizationID(org.ID).Save()
	token := jwks.JwtGenerator(t, member.LoginProviderID)

	// Fixture multiple environments
	envs := make(map[string]models.Environment)
	for i := 0; i < 3; i++ {
		env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()
		envs[env.ID] = env
	}

	resp, err := getClient(t, token).GetEnvironmentsWithResponse(ctx)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	for _, gotEnv := range resp.JSON200.Data {
		env, exists := envs[gotEnv.Id]
		require.True(t, exists)
		assert.Equal(t, env.Name, gotEnv.Name)
		assert.Equal(t, env.EnvType, string(gotEnv.Type))
		assert.Equal(t, env.OrganizationID, gotEnv.OrganizationId)
		assert.Equal(t, env.CreatedAt, gotEnv.CreatedAt)

		if env.ArchivedAt.Ptr() != nil && gotEnv.ArchivedAt != nil {
			tests.RequireEqualTime(t, env.ArchivedAt.Time, *gotEnv.ArchivedAt)
		}
	}
}
