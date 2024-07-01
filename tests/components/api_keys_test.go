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
	"github.com/subscribeddotdev/subscribed-backend/tests/jwks"
)

func TestApiKeys_Lifecycle(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()
	member := ff.NewMember().WithOrganizationID(org.ID).Save()
	token := jwks.JwtGenerator(t, member.LoginProviderID)

	t.Run("create_api_key", func(t *testing.T) {
		resp, err := getClient(t, token).CreateApiKey(
			ctx,
			&client.CreateApiKeyParams{EnvironmentId: env.ID},
			client.CreateApiKeyRequest{Name: gofakeit.AppName()},
		)
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})

	t.Run("create_api_key_with_expiration_date", func(t *testing.T) {
		resp, err := getClient(t, token).CreateApiKey(
			ctx,
			&client.CreateApiKeyParams{EnvironmentId: env.ID},
			client.CreateApiKeyRequest{Name: gofakeit.AppName(), ExpiresAt: toPtr(time.Now().Add(time.Hour * 24))},
		)
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})
}
