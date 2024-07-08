package command

import (
	"context"
	"time"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type CreateApiKey struct {
	OrgID         string
	Name          string
	EnvironmentID domain.EnvironmentID
	ExpiresAt     *time.Time
}

type CreateApiKeyHandler struct {
	apiKeyRepo domain.ApiKeyRepository
	envRepo    domain.EnvironmentRepository
}

func NewCreateApiKeyHandler(
	apiKeyRepo domain.ApiKeyRepository,
	envRepo domain.EnvironmentRepository,
) CreateApiKeyHandler {
	return CreateApiKeyHandler{
		apiKeyRepo: apiKeyRepo,
		envRepo:    envRepo,
	}
}

func (c CreateApiKeyHandler) Execute(ctx context.Context, cmd CreateApiKey) error {
	env, err := c.envRepo.ByID(ctx, cmd.EnvironmentID)
	if err != nil {
		return err
	}

	apiKey, err := domain.NewApiKey(cmd.Name, cmd.OrgID, cmd.EnvironmentID, cmd.ExpiresAt, env.Type() == domain.EnvTypeDevelopment)
	if err != nil {
		return err
	}

	return c.apiKeyRepo.Insert(ctx, apiKey)
}
