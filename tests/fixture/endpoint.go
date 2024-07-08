package fixture

import (
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Endpoint struct {
	factory      *Factory
	eventTypeIDs []string
	model        models.Endpoint
}

func (a *Endpoint) WithAppID(value string) *Endpoint {
	a.model.ApplicationID = value
	return a
}

func (a *Endpoint) WithEventTypeIDs(eventTypeIDs []string) *Endpoint {
	a.eventTypeIDs = eventTypeIDs
	return a
}

func (a *Endpoint) WithURL(value string) *Endpoint {
	a.model.URL = value
	return a
}

func (a *Endpoint) WithSigningSecret(value string) *Endpoint {
	a.model.SigningSecret = value
	return a
}

func (a *Endpoint) Save() models.Endpoint {
	err := a.model.Insert(a.factory.ctx, a.factory.db, boil.Infer())
	require.NoError(a.factory.t, err)

	if len(a.eventTypeIDs) > 0 {
		eventTypeIdModels := make([]*models.EventType, len(a.eventTypeIDs))
		for i, eventTypeID := range a.eventTypeIDs {
			eventTypeIdModels[i] = &models.EventType{
				ID: eventTypeID,
			}
		}
		err = a.model.AddEventTypes(a.factory.ctx, a.factory.db, false, eventTypeIdModels...)
		require.NoError(a.factory.t, err)
	}

	return a.model
}

func (a *Endpoint) NewDomainModel() *domain.Endpoint {
	endpointURL, err := domain.NewEndpointURL(a.model.URL)
	require.NoError(a.factory.t, err)

	endpoint, err := domain.NewEndpoint(
		endpointURL,
		domain.ApplicationID(a.model.ApplicationID),
		a.model.Description.String,
		nil, // TODO map event_type_ids correctly
	)
	require.NoError(a.factory.t, err)
	return endpoint
}
