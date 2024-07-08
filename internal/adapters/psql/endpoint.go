package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/friendsofgo/errors"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type EndpointRepository struct {
	db boil.ContextExecutor
}

func NewEndpointRepository(db boil.ContextExecutor) *EndpointRepository {
	return &EndpointRepository{
		db: db,
	}
}

func (o EndpointRepository) Insert(ctx context.Context, endpoint *domain.Endpoint) error {
	model := models.Endpoint{
		ID:            endpoint.ID().String(),
		ApplicationID: endpoint.ApplicationID().String(),
		URL:           endpoint.EndpointURL().String(),
		Description:   null.StringFrom(endpoint.Description()),
		SigningSecret: endpoint.SigningSecret().String(),
		CreatedAt:     endpoint.CreatedAt(),
		UpdatedAt:     endpoint.UpdatedAt(),
	}

	err := model.Insert(ctx, o.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("error saving endpoint: %v", err)
	}

	for _, eventTypeID := range endpoint.EventTypeIDs() {
		err = model.AddEventTypes(ctx, o.db, false, &models.EventType{
			ID: eventTypeID.String(),
		})
		if err != nil {
			return fmt.Errorf("error attaching endpoint to event_type with id '%s': %v", eventTypeID, err)
		}
	}

	return nil
}

func (o EndpointRepository) ByEventTypeIdAndAppID(
	ctx context.Context,
	eventTypeID domain.EventTypeID,
	appID domain.ApplicationID,
) ([]*domain.Endpoint, error) {
	model, err := models.EventTypes(
		models.EventTypeWhere.ID.EQ(eventTypeID.String()),
		qm.Load(models.EventTypeRels.Endpoints),
	).One(ctx, o.db)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrEventTypeNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("error querying event_types by id '%s': %v", eventTypeID, err)
	}

	return mapRowsToDomainEndpoints(model.R.Endpoints, appID)
}

func (o EndpointRepository) ByID(ctx context.Context, id domain.EndpointID) (*domain.Endpoint, error) {
	model, err := models.FindEndpoint(ctx, o.db, id.String())
	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrEndpointNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("error querying endpoint by id '%s': %v", id.String(), err)
	}

	return mapRowToDomainEndpoint(model)
}

func mapRowToDomainEndpoint(row *models.Endpoint) (*domain.Endpoint, error) {
	var eventTypeIDs []domain.EventTypeID

	if row.R != nil && row.R.EventTypes != nil {
		eventTypeIDs = make([]domain.EventTypeID, len(row.R.EventTypes))
		for j, eventType := range row.R.EventTypes {
			eventTypeIDs[j] = domain.EventTypeID(eventType.ID)
		}
	}

	endpoint, err := domain.UnMarshallEndpoint(
		domain.EndpointID(row.ID),
		domain.ApplicationID(row.ApplicationID),
		row.URL,
		row.Description.String,
		nil, // TODO: implement headers and then map it
		eventTypeIDs,
		row.SigningSecret,
		row.CreatedAt,
		row.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("error mapping row to endpoint: %v", err)
	}

	return endpoint, nil
}

func mapRowsToDomainEndpoints(rows []*models.Endpoint, appID domain.ApplicationID) ([]*domain.Endpoint, error) {
	endpoints := make([]*domain.Endpoint, len(rows))
	for i, row := range rows {
		if row.ApplicationID != appID.String() {
			continue
		}

		endpoint, err := mapRowToDomainEndpoint(row)
		if err != nil {
			return nil, err
		}

		endpoints[i] = endpoint
	}

	return endpoints, nil
}
