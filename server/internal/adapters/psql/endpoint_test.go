package psql_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed/server/internal/domain"
	"github.com/subscribeddotdev/subscribed/server/tests/fixture"
)

func TestEndpointRepository_Insert(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()
	app := ff.NewApplication().WithEnvironmentID(env.ID).Save()
	endpoint := fixtureEndpoint(t, domain.ApplicationID(app.ID), nil)

	require.NoError(t, endpointRepo.Insert(ctx, endpoint))
}

func TestEndpointRepository_ByEventTypeIdAndAppID(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()
	app := ff.NewApplication().WithEnvironmentID(env.ID).Save()
	eventType := ff.NewEventType().WithOrgID(org.ID).Save()
	eventTypeID := domain.EventTypeID(eventType.ID)

	fixtureEndpoints := make([]*domain.Endpoint, 5)
	for i := 0; i < 5; i++ {
		fixtureEndpoints[i] = fixtureEndpoint(t, domain.ApplicationID(app.ID), []domain.EventTypeID{eventTypeID})
		err := endpointRepo.Insert(ctx, fixtureEndpoints[i])
		require.NoError(t, err)
	}

	endpoints, err := endpointRepo.ByEventTypeIdAndAppID(ctx, eventTypeID, domain.ApplicationID(app.ID))
	require.NoError(t, err)
	assert.NotEmpty(t, endpoints)

	for _, endpoint := range endpoints {
		require.Equal(t, app.ID, endpoint.ApplicationID().String())
		containsEventTypeID := false

		for _, endpointEventTypeID := range endpoint.EventTypeIDs() {
			if endpointEventTypeID.String() == eventTypeID.String() {
				containsEventTypeID = true
				break
			}
		}

		assert.True(t, containsEventTypeID)
	}
}

func fixtureEndpoint(t *testing.T, appID domain.ApplicationID, eventTypeIDs []domain.EventTypeID) *domain.Endpoint {
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
