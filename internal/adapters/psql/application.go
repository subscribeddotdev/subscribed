package psql

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ApplicationRepository struct {
	db boil.ContextExecutor
}

func NewApplicationRepository(db boil.ContextExecutor) *ApplicationRepository {
	return &ApplicationRepository{
		db: db,
	}
}

func (o ApplicationRepository) Insert(ctx context.Context, application *domain.Application) error {
	model := models.Application{
		ID:            application.Id().String(),
		EnvironmentID: application.EnvID().String(),
		Name:          application.Name(),
		CreatedAt:     application.CreatedAt(),
	}

	err := model.Insert(ctx, o.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}
