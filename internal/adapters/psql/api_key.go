package psql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type ApiKeyRepository struct {
	db boil.ContextExecutor
}

func NewApiKeyRepository(db boil.ContextExecutor) *ApiKeyRepository {
	return &ApiKeyRepository{
		db: db,
	}
}

func (a ApiKeyRepository) Insert(ctx context.Context, apiKey *domain.ApiKey) error {
	model := models.APIKey{
		ID:            apiKey.Id().String(),
		SecretKey:     apiKey.SecretKey().FullKey(),
		Suffix:        apiKey.SecretKey().String(),
		OrgID:         apiKey.OrgID(),
		EnvironmentID: apiKey.EnvID().String(),
		Name:          apiKey.Name(),
		CreatedAt:     apiKey.CreatedAt(),
		ExpiresAt:     null.TimeFromPtr(apiKey.ExpiresAt()),
	}

	err := model.Insert(ctx, a.db, boil.Infer())
	var pqErr *pq.Error
	if err != nil {
		if ok := errors.As(err, &pqErr); ok {
			return domain.ErrApiKeyExists
		}

		return err
	}

	return nil
}

func (a ApiKeyRepository) FindBySecretKey(ctx context.Context, sk domain.SecretKey) (*domain.ApiKey, error) {
	model, err := models.APIKeys(models.APIKeyWhere.SecretKey.EQ(sk.FullKey())).One(ctx, a.db)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrApiKeyNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("error querying the api key '%s': %v", sk, err)
	}

	return domain.UnMarshallApiKey(
		domain.ApiKeyID(model.ID),
		domain.EnvironmentID(model.EnvironmentID),
		model.OrgID,
		model.Name,
		sk,
		model.CreatedAt,
		model.ExpiresAt.Ptr(),
	)
}

func (a ApiKeyRepository) FindByID(ctx context.Context, id domain.ApiKeyID) (*domain.ApiKey, error) {
	model, err := models.APIKeys(models.APIKeyWhere.ID.EQ(id.String())).One(ctx, a.db)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrApiKeyNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("error querying the api key: %v", err)
	}

	return domain.UnMarshallApiKey(
		domain.ApiKeyID(model.ID),
		domain.EnvironmentID(model.EnvironmentID),
		model.OrgID,
		model.Name,
		domain.SecretKey{}, // TODO Do we need another constructor without the secret key?
		model.CreatedAt,
		model.ExpiresAt.Ptr(),
	)
}

func (a ApiKeyRepository) FindAll(
	ctx context.Context,
	orgID string,
	envID domain.EnvironmentID,
) ([]*domain.ApiKey, error) {
	rows, err := models.APIKeys(
		models.APIKeyWhere.EnvironmentID.EQ(envID.String()),
		models.APIKeyWhere.OrgID.EQ(orgID),
	).All(ctx, a.db)
	if err != nil {
		return nil, err
	}

	apiKeys := make([]*domain.ApiKey, len(rows))
	for i, row := range rows {
		sk, err := domain.UnMarshallSecretKey(row.SecretKey)
		if err != nil {
			return nil, err
		}

		ak, err := domain.UnMarshallApiKey(
			domain.ApiKeyID(row.ID),
			envID,
			orgID,
			row.Name,
			sk,
			row.CreatedAt,
			row.ExpiresAt.Ptr(),
		)
		if err != nil {
			return nil, err
		}

		apiKeys[i] = ak
	}

	return apiKeys, nil
}

func (a ApiKeyRepository) Destroy(ctx context.Context, orgID string, id domain.ApiKeyID) error {
	row, err := models.APIKeys(models.APIKeyWhere.OrgID.EQ(orgID), models.APIKeyWhere.ID.EQ(id.String())).One(ctx, a.db)
	if errors.Is(err, sql.ErrNoRows) {
		return domain.ErrApiKeyNotFound
	}
	if err != nil {
		return fmt.Errorf("error querying api key via id '%s': %v", id, err)
	}

	_, err = row.Delete(ctx, a.db)
	if err != nil {
		return fmt.Errorf("error desotrying api key with id '%s': %v", id, err)
	}

	return nil
}
