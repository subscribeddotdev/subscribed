package query

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type AllApplications struct {
	PaginationParams
	EnvironmentID string
	OrgID         string
}

type allApplicationsHandler struct {
}

func NewAllApplicationsHandler() allApplicationsHandler {
	return allApplicationsHandler{}
}

func (h allApplicationsHandler) Execute(ctx context.Context, q AllApplications) (Paginated[[]domain.Application], error) {
	return Paginated[[]domain.Application]{}, nil
}
