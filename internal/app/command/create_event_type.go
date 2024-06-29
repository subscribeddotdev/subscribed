package command

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type CreateEventType struct {
	OrgID         string
	Name          string
	Description   string
	Schema        string
	SchemaExample string
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
	orgID, err := domain.NewIdFromString(cmd.OrgID)
	if err != nil {
		return err
	}

	eventType, err := domain.NewEventType(orgID, cmd.Name, cmd.Description, cmd.Schema, cmd.SchemaExample)
	if err != nil {
		return err
	}

	return c.repo.Insert(ctx, eventType)
}
