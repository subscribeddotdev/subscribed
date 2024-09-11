package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/friendsofgo/errors"
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
	total, err := models.EventTypes(models.EventTypeWhere.OrgID.EQ(orgID.String())).Count(ctx, e.db)
	if err != nil {
		return query.Paginated[[]domain.EventType]{}, fmt.Errorf("error counting event types: %v", err)
	}

	rows, err := models.EventTypes(
		models.EventTypeWhere.OrgID.EQ(orgID.String()),
		qm.Offset(mapPaginationParamsToSqlOffset(pagination)),
		qm.Limit(pagination.Limit()),
		qm.OrderBy("created_at DESC"),
	).All(ctx, e.db)
	if err != nil {
		return query.Paginated[[]domain.EventType]{}, fmt.Errorf("error querying event types: %v", err)
	}

	return query.Paginated[[]domain.EventType]{
		Total:       int(total),
		PerPage:     len(rows),
		CurrentPage: pagination.Page(),
		TotalPages:  getPaginationTotalPages(total, pagination.Limit()),
		Data:        mapRowsToEventTypes(rows),
	}, nil
}

func (e EventTypeRepository) ByID(ctx context.Context, orgID string, id domain.EventTypeID) (*domain.EventType, error) {
	row, err := models.EventTypes(
		models.EventTypeWhere.OrgID.EQ(orgID),
		models.EventTypeWhere.ID.EQ(id.String()),
	).One(ctx, e.db)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrEventTypeNotFound
	}

	if err != nil {
		return nil, fmt.Errorf("error querying event type with id '%s': %v", id, err)
	}

	return mapRowToEventType(row), nil
}

func mapRowsToEventTypes(rows []*models.EventType) []domain.EventType {
	eventTypes := make([]domain.EventType, len(rows))
	for i, row := range rows {
		eventType := mapRowToEventType(row)
		eventTypes[i] = *eventType
	}

	return eventTypes
}

func mapRowToEventType(row *models.EventType) *domain.EventType {
	return domain.UnMarshallEventType(
		domain.EventTypeID(row.ID),
		row.OrgID,
		row.Name,
		row.Description.String,
		row.Schema.String,
		row.SchemaExample.String,
		row.CreatedAt,
		row.ArchivedAt.Ptr(),
	)
}
