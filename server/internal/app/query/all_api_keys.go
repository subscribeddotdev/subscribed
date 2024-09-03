package query

import (
	"context"

	"github.com/subscribeddotdev/subscribed/server/internal/domain"
)

type AllApiKeys struct {
	OrgID         string
	EnvironmentID string
}

type allApiKeysHandler struct {
	repo domain.ApiKeyRepository
}

func NewAllApiKeysHandler(repo domain.ApiKeyRepository) allApiKeysHandler {
	return allApiKeysHandler{repo: repo}
}

func (h allApiKeysHandler) Execute(ctx context.Context, q AllApiKeys) ([]*domain.ApiKey, error) {
	return h.repo.FindAll(ctx, q.OrgID, domain.EnvironmentID(q.EnvironmentID))
}
