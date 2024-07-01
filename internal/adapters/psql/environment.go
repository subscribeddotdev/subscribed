package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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

func (o EnvironmentRepository) ByID(ctx context.Context, id domain.ID) (*domain.Environment, error) {
	model, err := models.FindEnvironment(ctx, o.db, id.String())
	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrEnvironmentNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("error querying the environment by its id '%s': %v", id, err)
	}

	return domain.UnMarshallEnvironment(
		model.ID,
		model.OrganizationID,
		model.Name,
		model.EnvType,
		model.CreatedAt,
		model.ArchivedAt.Ptr(),
	)
}
