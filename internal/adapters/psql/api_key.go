package psql

import (
	"context"
	"errors"

	"github.com/lib/pq"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ApiKeyRepository struct {
	db boil.ContextExecutor
}

func NewApiKeyRepository(db boil.ContextExecutor) *ApiKeyRepository {
	return &ApiKeyRepository{
		db: db,
	}
}

func (o ApiKeyRepository) Insert(ctx context.Context, apiKey *domain.ApiKey) error {
	model := models.APIKey{
		SecretKey:     apiKey.SecretKey().FullKey(),
		Suffix:        apiKey.SecretKey().String(),
		EnvironmentID: apiKey.EnvID().String(),
		Name:          apiKey.Name(),
		CreatedAt:     apiKey.CreatedAt(),
		ExpiresAt:     null.TimeFromPtr(apiKey.ExpiresAt()),
	}

	err := model.Insert(ctx, o.db, boil.Infer())
	var pqErr *pq.Error
	if err != nil {
		if ok := errors.As(err, &pqErr); ok {
			return domain.ErrApiKeyExists
		}

		return err
	}

	return nil
}
