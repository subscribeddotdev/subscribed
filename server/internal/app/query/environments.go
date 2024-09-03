package query

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type AllEnvironments struct {
	OrgID string
}

type allEnvironmentsHandler struct {
	envRepo domain.EnvironmentRepository
}

func NewEnvironmentsHandler(envRepo domain.EnvironmentRepository) allEnvironmentsHandler {
	return allEnvironmentsHandler{
		envRepo: envRepo,
	}
}

func (h allEnvironmentsHandler) Execute(ctx context.Context, q AllEnvironments) ([]*domain.Environment, error) {
	return h.envRepo.FindAll(ctx, q.OrgID)
}
