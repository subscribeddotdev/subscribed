package psql

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type EventTypeRepository struct {
	db boil.ContextExecutor
}

func NewEventTypeRepository(db boil.ContextExecutor) *EventTypeRepository {
	return &EventTypeRepository{
		db: db,
	}
}

func (o EventTypeRepository) Insert(ctx context.Context, eventType *domain.EventType) error
	return nil
}
