package psql_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/tests"
	"github.com/subscribeddotdev/subscribed-backend/tests/fixture"
)

func TestApiKeyRepository_Insert(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()

	apiKey, err := domain.NewApiKey(gofakeit.AppName(), tests.MustID(t, env.ID), nil, false)
	require.NoError(t, err)
	require.NoError(t, apiKeyRepo.Insert(ctx, apiKey))

	t.Run("error_api_key_already_exists", func(t *testing.T) {
		assert.ErrorIs(t, apiKeyRepo.Insert(ctx, apiKey), domain.ErrApiKeyExists)
	})
}
