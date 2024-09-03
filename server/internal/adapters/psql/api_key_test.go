package psql_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed/server/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed/server/internal/domain"
	"github.com/subscribeddotdev/subscribed/server/tests/fixture"
)

func TestApiKeyRepository_Insert(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()

	apiKey, err := domain.NewApiKey(
		gofakeit.AppName(),
		org.ID,
		domain.EnvironmentID(env.ID),
		nil,
		false,
	)
	require.NoError(t, err)
	require.NoError(t, apiKeyRepo.Insert(ctx, apiKey))

	t.Run("error_api_key_already_exists", func(t *testing.T) {
		assert.ErrorIs(t, apiKeyRepo.Insert(ctx, apiKey), domain.ErrApiKeyExists)
	})
}

func TestApiKeyRepository_Destroy(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()
	apiKey := ff.NewApiKey().WithOrgID(org.ID).WithEnvironmentID(env.ID).Save()

	err := apiKeyRepo.Destroy(ctx, org.ID, domain.ApiKeyID(apiKey.ID))
	require.NoError(t, err)

	err = apiKeyRepo.Destroy(ctx, org.ID, domain.ApiKeyID(apiKey.ID))
	require.ErrorIs(t, err, domain.ErrApiKeyNotFound)
}

func TestApiKeyRepository_FindAll(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	env1 := ff.NewEnvironment().WithOrganizationID(org.ID).Save()
	env2 := ff.NewEnvironment().WithOrganizationID(org.ID).Save()

	fixtureApiKeys := make(map[string]models.APIKey)
	for i := 0; i < 10; i++ {
		var ak models.APIKey
		if i%2 == 0 {
			ak = ff.NewApiKey().WithEnvironmentID(env1.ID).WithOrgID(org.ID).Save()
		} else {
			ak = ff.NewApiKey().WithEnvironmentID(env2.ID).WithOrgID(org.ID).Save()
		}
		fixtureApiKeys[ak.SecretKey] = ak
	}

	apiKeys, err := apiKeyRepo.FindAll(ctx, org.ID, domain.EnvironmentID(env1.ID))
	require.NoError(t, err)
	require.NotEmpty(t, apiKeys)

	for _, apiKey := range apiKeys {
		_, exists := fixtureApiKeys[apiKey.SecretKey().FullKey()]
		require.True(t, exists)

		require.Equal(t, org.ID, apiKey.OrgID(), "api key must belong to the target organisation")
		require.Equal(t, env1.ID, apiKey.EnvID().String(), "api key must belong to the target environment")
	}
}
