package command

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

type CreateEventType struct {
	OrgID         iam.OrgID
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
	eventType, err := domain.NewEventType(cmd.OrgID.String(), cmd.Name, cmd.Description, cmd.Schema, cmd.SchemaExample)
	if err != nil {
		return err
	}

	return c.repo.Insert(ctx, eventType)
}
