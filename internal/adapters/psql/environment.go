package psql

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type EnvironmentRepository struct {
	db boil.ContextExecutor
}

func NewEnvironmentRepository(db boil.ContextExecutor) *EnvironmentRepository {
	return &EnvironmentRepository{
		db: db,
	}
}

func (o EnvironmentRepository) Insert(ctx context.Context, env *domain.Environment) error {
	model := models.Environment{
		ID:             env.Id().String(),
		OrganizationID: env.OrgID().String(),
		Name:           env.Name(),
		EnvType:        env.Type().String(),
		CreatedAt:      env.CreatedAt(),
		ArchivedAt:     null.TimeFromPtr(env.ArchivedAt()),
	}

	err := model.Insert(ctx, o.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}
