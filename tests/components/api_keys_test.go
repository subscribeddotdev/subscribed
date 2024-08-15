package components_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/tests/client"
	"github.com/subscribeddotdev/subscribed-backend/tests/fixture"
)

func TestApiKeys_Lifecycle(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()
	member, password := ff.NewMember().WithOrganizationID(org.ID).Save()
	token := signIn(t, member.Email, password)
	apiClient := getClient(t, token)

	t.Run("create_api_key", func(t *testing.T) {
		resp, err := apiClient.CreateApiKey(
			ctx,
			client.CreateApiKeyRequest{
				Name:          gofakeit.AppName(),
				EnvironmentId: env.ID,
			},
		)
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})

	t.Run("create_api_key_with_expiration_date", func(t *testing.T) {
		resp, err := apiClient.CreateApiKey(
			ctx,
			client.CreateApiKeyRequest{
				Name:          gofakeit.AppName(),
				ExpiresAt:     toPtr(time.Now().Add(time.Hour * 24)),
				EnvironmentId: env.ID,
			},
		)
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})

	t.Run("get_all_api_keys_by_environment", func(t *testing.T) {
		for i := 0; i < 10; i++ {
			ff.NewApiKey().WithEnvironmentID(env.ID).WithOrgID(org.ID).Save()
		}

		resp, err := apiClient.GetAllApiKeysWithResponse(ctx, &client.GetAllApiKeysParams{EnvironmentId: env.ID})
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode())
		require.NotEmpty(t, resp.JSON200.Data)
	})
}
