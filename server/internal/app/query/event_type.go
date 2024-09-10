package query

import (
	"context"

	"github.com/subscribeddotdev/subscribed/server/internal/domain"
)

type EventType struct {
	OrgID       string
	EventTypeID string
}

type eventTypeHandler struct {
	repo domain.EventTypeRepository
}

func NewEventTypeHandler(repo domain.EventTypeRepository) eventTypeHandler {
	return eventTypeHandler{repo: repo}
}

func (h eventTypeHandler) Execute(ctx context.Context, q EventType) (*domain.EventType, error) {
	return nil, nil
}
