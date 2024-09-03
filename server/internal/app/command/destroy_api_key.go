package command

import (
	"context"

	"github.com/subscribeddotdev/subscribed/server/internal/domain"
)

type DestroyApiKey struct {
	OrgID string
	ID    domain.ApiKeyID
}

type DestroyApiKeyHandler struct {
	apiKeyRepo domain.ApiKeyRepository
}

func NewDestroyApiKeyHandler(
	apiKeyRepo domain.ApiKeyRepository,
) DestroyApiKeyHandler {
	return DestroyApiKeyHandler{
		apiKeyRepo: apiKeyRepo,
	}
}

func (c DestroyApiKeyHandler) Execute(ctx context.Context, cmd DestroyApiKey) error {
	return c.apiKeyRepo.Destroy(ctx, cmd.OrgID, cmd.ID)
}
