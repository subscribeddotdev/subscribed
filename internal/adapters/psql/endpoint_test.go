package psql_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

func TestEndpointRepository_Insert(t *testing.T) {
	org := fixtureAndSaveOrganization(t)
	env := fixtureEnvironment(t, org.ID)
	app := fixtureApplication(t, env.Id())
	endpoint := fixtureEndpoint(t, app.Id())

	require.NoError(t, environmentRepo.Insert(ctx, env))
	require.NoError(t, applicationRepo.Insert(ctx, app))
	require.NoError(t, endpointRepo.Insert(ctx, endpoint))
}

func fixtureEndpoint(t *testing.T, appID domain.ID) *domain.Endpoint {
	endpointURL, err := domain.NewEndpointURL(gofakeit.URL())
	require.NoError(t, err)

	endpoint, err := domain.NewEndpoint(endpointURL, appID, gofakeit.Sentence(5), nil)
	require.NoError(t, err)

	return endpoint
}
