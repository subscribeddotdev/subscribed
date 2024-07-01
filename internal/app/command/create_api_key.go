package command

import (
	"context"
	"time"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type CreateApiKey struct {
	Name          string
	EnvironmentID string
	ExpiresAt     *time.Time
}

type CreateApiKeyHandler struct {
	apiKeyRepo domain.ApiKeyRepository
	envRepo    domain.EnvironmentRepository
}

func NewCreateApiKeyHandler(apiKeyRepo domain.ApiKeyRepository, envRepo domain.EnvironmentRepository) CreateApiKeyHandler {
	return CreateApiKeyHandler{
		apiKeyRepo: apiKeyRepo,
		envRepo:    envRepo,
	}
}

func (c CreateApiKeyHandler) Execute(ctx context.Context, cmd CreateApiKey) error {
	envID, err := domain.NewIdFromString(cmd.EnvironmentID)
	if err != nil {
		return err
	}

	env, err := c.envRepo.ByID(ctx, envID)
	if err != nil {
		return err
	}

	apiKey, err := domain.NewApiKey(cmd.Name, envID, cmd.ExpiresAt, env.Type() == domain.EnvTypeDevelopment)
	if err != nil {
		return err
	}

	return c.apiKeyRepo.Insert(ctx, apiKey)
}
