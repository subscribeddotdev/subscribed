package command

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type CreateEventType struct {
}

type CreateEventTypeHandler struct {
	repo domain.EventTypeRepository
}

func NewCreateEventTypeHandler(repo domain.EventTypeRepository) CreateEventTypeHandler {
	return CreateEventTypeHandler{
		repo: repo,
	}
}

func (c CreateEventTypeHandler) Execute(ctx context.Context, cmd CreateEventType) error {
	return nil
}
