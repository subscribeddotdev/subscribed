package psql

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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
		ID:            endpoint.Id().String(),
		ApplicationID: endpoint.ApplicationID().String(),
		URL:           endpoint.EndpointURL().String(),
		Description:   null.StringFrom(endpoint.Description()),
		SigningSecret: endpoint.SigningSecret().String(),
		CreatedAt:     endpoint.CreatedAt(),
		UpdatedAt:     endpoint.UpdatedAt(),
	}

	err := model.Insert(ctx, o.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}
