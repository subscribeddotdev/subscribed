package psql

import (
	"context"
	"fmt"

	"github.com/subscribeddotdev/subscribed/server/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed/server/internal/app/query"
	"github.com/subscribeddotdev/subscribed/server/internal/domain"
	"github.com/subscribeddotdev/subscribed/server/internal/domain/iam"
	"github.com/volatiletech/null/v8"
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

func (e EventTypeRepository) Insert(ctx context.Context, eventType *domain.EventType) error {
	model := models.EventType{
		ID:            eventType.ID().String(),
		OrgID:         eventType.OrgID(),
		Name:          eventType.Name(),
		Schema:        null.StringFrom(eventType.Schema()),
		Description:   null.StringFrom(eventType.Description()),
		SchemaExample: null.StringFrom(eventType.SchemaExample()),
		ArchivedAt:    null.TimeFromPtr(eventType.ArchivedAt()),
		CreatedAt:     eventType.CreatedAt(),
	}

	err := model.Insert(ctx, e.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("unable to save event type: %v", err)
	}

	return nil
}

func (e EventTypeRepository) FindAll(
	ctx context.Context,
	envID domain.EnvironmentID,
	orgID iam.OrgID,
	pagination query.PaginationParams,
) (query.Paginated[[]domain.EventType], error) {
	return query.Paginated[[]domain.EventType]{}, nil
}
