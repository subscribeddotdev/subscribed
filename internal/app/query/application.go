package query

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type Application struct {
	OrgID         string
	ApplicationID string
}

type applicationHandler struct {
	applicationsFinder applicationsFinder
}

func NewApplicationHandler(applicationsFinder applicationsFinder) applicationHandler {
	return applicationHandler{applicationsFinder: applicationsFinder}
}

func (h applicationHandler) Execute(ctx context.Context, q Application) (*domain.Application, error) {
	return nil, nil
}
