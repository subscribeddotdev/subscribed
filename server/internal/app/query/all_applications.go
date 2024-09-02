package query

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

type AllApplications struct {
	PaginationParams
	EnvironmentID string
	OrgID         string
}

type applicationsFinder interface {
	FindAll(
		ctx context.Context,
		envID domain.EnvironmentID,
		orgID iam.OrgID,
		pagination PaginationParams,
	) (Paginated[[]domain.Application], error)
	FindByID(ctx context.Context, id domain.ApplicationID, orgID iam.OrgID) (*domain.Application, error)
}

type allApplicationsHandler struct {
	applicationsFinder applicationsFinder
}

func NewAllApplicationsHandler(applicationsFinder applicationsFinder) allApplicationsHandler {
	return allApplicationsHandler{
		applicationsFinder: applicationsFinder,
	}
}

func (h allApplicationsHandler) Execute(ctx context.Context, q AllApplications) (Paginated[[]domain.Application], error) {
	return h.applicationsFinder.FindAll(
		ctx,
		domain.EnvironmentID(q.EnvironmentID),
		iam.OrgID(q.OrgID),
		q.PaginationParams,
	)
}
