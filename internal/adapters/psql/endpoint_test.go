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

func TestEndpointRepository_Insert(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()
	app := ff.NewApplication().WithEnvironmentID(env.ID).Save()
	endpoint := fixtureEndpoint(t, tests.MustID(t, app.ID), nil)

	require.NoError(t, endpointRepo.Insert(ctx, endpoint))
}

func TestEndpointRepository_ByEventTypeIdAndAppID(t *testing.T) {
	// Fixture several
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()
	app := ff.NewApplication().WithEnvironmentID(env.ID).Save()
	eventType := ff.NewEventType().WithOrgID(org.ID).Save()
	eventTypeID := tests.MustID(t, app.ID)

	endpoints := make([]*domain.Endpoint, 5)
	for i := 0; i < 5; i++ {
		endpoints[i] = fixtureEndpoint(t, eventTypeID, []domain.ID{tests.MustID(t, eventType.ID)})
		err := endpointRepo.Insert(ctx, endpoints[i])
		require.NoError(t, err)
	}

	result, err := endpointRepo.ByEventTypeIdAndAppID(ctx, eventTypeID, tests.MustID(t, app.ID))
	require.NoError(t, err)
	assert.NotEmpty(t, result)
}

func fixtureEndpoint(t *testing.T, appID domain.ID, eventTypeIDs []domain.ID) *domain.Endpoint {
	endpointURL, err := domain.NewEndpointURL(gofakeit.URL())
	require.NoError(t, err)

	endpoint, err := domain.NewEndpoint(
		endpointURL,
		appID,
		gofakeit.Sentence(5),
		eventTypeIDs,
	)
	require.NoError(t, err)

	return endpoint
}
