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
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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
	orgID iam.OrgID,
	pagination query.PaginationParams,
) (query.Paginated[[]domain.EventType], error) {
	total, err := models.Applications(
		//TODO: include orgID in the query
		models.ApplicationWhere.EnvironmentID.EQ(envID.String()),
	).Count(ctx, e.db)
	if err != nil {
		return query.Paginated[[]domain.EventType]{}, fmt.Errorf("error counting applications: %v", err)
	}

	rows, err := models.EventTypes(
		//TODO: include orgID in the query
		// models.EventTypeWhere.OrgID.EQ(orgID),
		models.ApplicationWhere.EnvironmentID.EQ(envID.String()),
		qm.Offset(mapPaginationParamsToSqlOffset(pagination)),
		qm.Limit(pagination.Limit()),
		qm.OrderBy("created_at DESC"),
	).All(ctx, e.db)
	if err != nil {
		return query.Paginated[[]domain.EventType]{}, fmt.Errorf("error querying applications: %v", err)
	}

	return query.Paginated[[]domain.EventType]{
		Total:       int(total),
		PerPage:     len(rows),
		CurrentPage: pagination.Page(),
		TotalPages:  getPaginationTotalPages(total, pagination.Limit()),
		Data:        mapRowsToEventTypes(rows),
	}, nil

	return query.Paginated[[]domain.EventType]{}, nil
}

func mapRowsToEventTypes(rows []*models.EventType) []domain.EventType {
	return nil
}
