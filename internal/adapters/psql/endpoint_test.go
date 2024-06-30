package psql_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/tests"
	"github.com/subscribeddotdev/subscribed-backend/tests/fixture"
)

func TestEndpointRepository_Insert(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()
	app := ff.NewApplication().WithEnvironmentID(env.ID).Save()
	endpoint := fixtureEndpoint(t, tests.MustID(t, app.ID))

	require.NoError(t, endpointRepo.Insert(ctx, endpoint))
}

func fixtureEndpoint(t *testing.T, appID domain.ID) *domain.Endpoint {
	endpointURL, err := domain.NewEndpointURL(gofakeit.URL())
	require.NoError(t, err)

	endpoint, err := domain.NewEndpoint(
		endpointURL,
		appID,
		gofakeit.Sentence(5),
		nil,
	)
	require.NoError(t, err)

	return endpoint
}
