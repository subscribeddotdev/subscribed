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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ApplicationRepository struct {
	db boil.ContextExecutor
}

func NewApplicationRepository(db boil.ContextExecutor) *ApplicationRepository {
	return &ApplicationRepository{
		db: db,
	}
}

func (o ApplicationRepository) Insert(ctx context.Context, application *domain.Application) error {
	model := models.Application{
		ID:            application.ID().String(),
		EnvironmentID: application.EnvID().String(),
		Name:          application.Name(),
		CreatedAt:     application.CreatedAt(),
	}

	err := model.Insert(ctx, o.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (o ApplicationRepository) FindAll(
	ctx context.Context,
	envID domain.EnvironmentID,
	orgID iam.OrgID, // TODO: to be used later
	pagination query.PaginationParams,
) (query.Paginated[[]domain.Application], error) {
	total, err := models.Applications(
		//TODO: include orgID in the query
		models.ApplicationWhere.EnvironmentID.EQ(envID.String()),
	).Count(ctx, o.db)
	if err != nil {
		return query.Paginated[[]domain.Application]{}, fmt.Errorf("error counting applications: %v", err)
	}

	rows, err := models.Applications(
		//TODO: include orgID in the query
		models.ApplicationWhere.EnvironmentID.EQ(envID.String()),
		qm.Offset(mapPaginationParamsToSqlOffset(pagination)),
		qm.Limit(pagination.Limit()),
		qm.OrderBy("created_at DESC"),
	).All(ctx, o.db)
	if err != nil {
		return query.Paginated[[]domain.Application]{}, fmt.Errorf("error querying applications: %v", err)
	}

	return query.Paginated[[]domain.Application]{
		Total:       int(total),
		PerPage:     len(rows),
		CurrentPage: pagination.Page(),
		TotalPages:  getPaginationTotalPages(total, pagination.Limit()),
		Data:        mapRowsToApplications(rows),
	}, nil
}

func (o ApplicationRepository) FindByID(
	ctx context.Context,
	id domain.ApplicationID,
	orgID iam.OrgID,
) (*domain.Application, error) {
	row, err := models.Applications(
		models.ApplicationWhere.ID.EQ(id.String()),
		// TODO: models.ApplicationWhere.OrgID.EQ(orgID.String())
	).One(ctx, o.db)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrAppNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("error querying application via id: %v", err)
	}

	return domain.UnMarshallApplication(
		domain.ApplicationID(row.ID),
		row.Name,
		domain.EnvironmentID(row.EnvironmentID),
		// TODO: iam.OrgID(row.OrgID),
		row.CreatedAt,
	), nil
}

func mapPaginationParamsToSqlOffset(pagination query.PaginationParams) int {
	if pagination.Page() == 1 {
		return 0
	}

	// Page is 1-based index whereas an SQL offset is 0-based index hence the subtraction
	return (pagination.Page() - 1) * pagination.Limit()
}

func getPaginationTotalPages(total int64, limit int) int {
	return int(total / int64(limit))
}

func mapRowsToApplications(rows []*models.Application) []domain.Application {
	apps := make([]domain.Application, len(rows))
	for i, row := range rows {
		app := domain.UnMarshallApplication(
			domain.ApplicationID(row.ID),
			row.Name,
			domain.EnvironmentID(rows[i].EnvironmentID),
			row.CreatedAt,
		)

		apps[i] = *app
	}

	return apps
}
