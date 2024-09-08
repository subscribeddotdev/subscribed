package psql_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed/server/internal/app/query"
	"github.com/subscribeddotdev/subscribed/server/internal/domain"
	"github.com/subscribeddotdev/subscribed/server/internal/domain/iam"
	"github.com/subscribeddotdev/subscribed/server/tests"
	"github.com/subscribeddotdev/subscribed/server/tests/fixture"
)

func TestEventTypesRepository_FindAll(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()

	eventTypesFixtureCount := 20
	for i := 0; i < eventTypesFixtureCount; i++ {
		ff.NewEventType().WithOrgID(org.ID).Save()
	}

	t.Run("return_an_empty_slice_when_page_is_out_of_bounds", func(t *testing.T) {
		result, err := eventTypeRepo.FindAll(
			ctx,
			iam.OrgID(org.ID),
			query.NewPaginationParams(tests.ToPtr(30), tests.ToPtr(5)),
		)
		require.NoError(t, err)
		require.Empty(t, result.Data)
		require.Equal(t, 0, result.PerPage)
	})

	t.Run("iteratively_query_event_types_from_different_pages", func(t *testing.T) {
		perPage := 5
		totalPages := eventTypesFixtureCount / 5
		queriedAppIDs := make(map[string]domain.EventType)

		for i := 0; i < totalPages; i++ {
			currentPage := i + 1
			result, err := eventTypeRepo.FindAll(ctx, iam.OrgID(org.ID), query.NewPaginationParams(&currentPage, &perPage))
			require.NoError(t, err)
			require.NotEmpty(t, result.Data)
			require.Equal(t, perPage, result.PerPage)
			require.Equal(t, currentPage, result.CurrentPage)
			require.Equal(t, eventTypesFixtureCount, result.Total)
			require.Equal(t, totalPages, result.TotalPages)

			for _, app := range result.Data {
				_, exists := queriedAppIDs[app.ID().String()]
				require.Falsef(t, exists, "event type must not have already been returned: app id '%s'", app.ID())
				queriedAppIDs[app.ID().String()] = app
			}
		}
	})
}
