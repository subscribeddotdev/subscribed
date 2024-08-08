package query

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type Environments struct {
	OrgID string
}

type environmentsHandler struct {
	envRepo domain.EnvironmentRepository
}

func NewEnvironmentsHandler(envRepo domain.EnvironmentRepository) environmentsHandler {
	return environmentsHandler{envRepo: envRepo}
}

func (h environmentsHandler) Execute(ctx context.Context, q Environments) ([]*domain.Environment, error) {
	return h.envRepo.FindAll(ctx, q.OrgID)
}
