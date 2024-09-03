package psql

import (
	"context"

	"github.com/subscribeddotdev/subscribed/server/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed/server/internal/domain/iam"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type OrganizationRepository struct {
	db boil.ContextExecutor
}

func NewOrganizationRepository(db boil.ContextExecutor) *OrganizationRepository {
	return &OrganizationRepository{
		db: db,
	}
}

func (o OrganizationRepository) Insert(ctx context.Context, org *iam.Organization) error {
	model := models.Organization{
		ID:         org.ID().String(),
		CreatedAt:  org.CreatedAt(),
		DisabledAt: null.TimeFromPtr(org.DisabledAt()),
	}

	err := model.Insert(ctx, o.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}
