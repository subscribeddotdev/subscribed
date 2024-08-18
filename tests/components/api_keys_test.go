package components_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/tests"
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
		reqBody := client.CreateApiKeyRequest{
			Name:          gofakeit.AppName(),
			EnvironmentId: env.ID,
		}
		resp, err := apiClient.CreateApiKeyWithResponse(
			ctx,
			reqBody,
		)
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, resp.StatusCode())
		require.NotEmpty(t, resp.JSON201.UnmaskedApiKey)

		savedApiKey := findApiKey(t, reqBody.Name, reqBody.EnvironmentId)
		require.Equal(t, savedApiKey.SecretKey, resp.JSON201.UnmaskedApiKey)
	})

	t.Run("create_api_key_with_expiration_date", func(t *testing.T) {
		reqBody := client.CreateApiKeyRequest{
			Name:          gofakeit.AppName(),
			ExpiresAt:     toPtr(time.Now().Add(time.Hour * 24)),
			EnvironmentId: env.ID,
		}
		resp, err := apiClient.CreateApiKeyWithResponse(
			ctx,
			reqBody,
		)
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, resp.StatusCode())
		require.NotEmpty(t, resp.JSON201.UnmaskedApiKey)

		savedApiKey := findApiKey(t, reqBody.Name, reqBody.EnvironmentId)
		require.Equal(t, savedApiKey.SecretKey, resp.JSON201.UnmaskedApiKey)
		tests.RequireEqualTime(t, *reqBody.ExpiresAt, savedApiKey.ExpiresAt.Time)
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

	t.Run("destroy_an_api_key_by_id", func(t *testing.T) {
		apiKey := ff.NewApiKey().WithOrgID(org.ID).WithEnvironmentID(env.ID).Save()
		resp, err := apiClient.DestroyApiKey(ctx, apiKey.ID)
		require.NoError(t, err)
		require.Equal(t, http.StatusNoContent, resp.StatusCode)

		// Attempt to destroy an already destroyed api key
		resp, err = apiClient.DestroyApiKey(ctx, apiKey.ID)
		require.NoError(t, err)
		require.Equal(t, http.StatusNotFound, resp.StatusCode)
	})
}

func findApiKey(t *testing.T, name, environmentId string) *models.APIKey {
	row, err := models.APIKeys(
		models.APIKeyWhere.Name.EQ(name),
		models.APIKeyWhere.EnvironmentID.EQ(environmentId),
	).One(ctx, db)
	require.NoError(t, err)
	return row
}
