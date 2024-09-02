package psql_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/app/query"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
	"github.com/subscribeddotdev/subscribed-backend/tests"
	"github.com/subscribeddotdev/subscribed-backend/tests/fixture"
)

func TestApplicationRepository_Insert(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()
	app := ff.NewApplication().WithEnvironmentID(env.ID).NewDomainModel()

	require.NoError(t, applicationRepo.Insert(ctx, app))
}

func TestApplicationRepository_FindAll(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()

	applicationFixtureCount := 20
	for i := 0; i < applicationFixtureCount; i++ {
		ff.NewApplication().WithEnvironmentID(env.ID).Save()
	}

	t.Run("return_an_empty_slice_when_page_is_out_of_bounds", func(t *testing.T) {
		result, err := applicationRepo.FindAll(
			ctx,
			domain.EnvironmentID(env.ID),
			iam.OrgID(env.ID),
			query.NewPaginationParams(tests.ToPtr(30), tests.ToPtr(5)),
		)
		require.NoError(t, err)
		require.Empty(t, result.Data)
		require.Equal(t, 0, result.PerPage)
	})

	t.Run("iteratively_query_applications_from_different_pages", func(t *testing.T) {
		perPage := 5
		totalPages := applicationFixtureCount / 5
		queriedAppIDs := make(map[string]domain.Application)

		for i := 0; i < totalPages; i++ {
			currentPage := i + 1
			result, err := applicationRepo.FindAll(ctx, domain.EnvironmentID(env.ID), iam.OrgID(env.ID), query.NewPaginationParams(&currentPage, &perPage))
			require.NoError(t, err)
			require.NotEmpty(t, result.Data)
			require.Equal(t, perPage, result.PerPage)
			require.Equal(t, currentPage, result.CurrentPage)
			require.Equal(t, applicationFixtureCount, result.Total)
			require.Equal(t, totalPages, result.TotalPages)

			for _, app := range result.Data {
				_, exists := queriedAppIDs[app.ID().String()]
				require.Falsef(t, exists, "application must not have already been returned: app id '%s'", app.ID())
				queriedAppIDs[app.ID().String()] = app
			}
		}
	})
}

func TestApplicationRepository_FindById(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()

	t.Run("find_app_by_id_and_org_id", func(t *testing.T) {
		appFixture := ff.NewApplication().WithEnvironmentID(env.ID).Save()
		app, err := applicationRepo.FindByID(ctx, domain.ApplicationID(appFixture.ID), iam.OrgID(org.ID))

		require.NoError(t, err)
		require.NotNil(t, app)
		require.Equal(t, appFixture.ID, app.ID().String())
	})

	t.Run("error_app_not_found_wrong_id", func(t *testing.T) {
		app, err := applicationRepo.FindByID(ctx, domain.NewApplicationID(), iam.OrgID(org.ID))
		require.ErrorIs(t, err, domain.ErrAppNotFound)
		require.Nil(t, app)
	})

	//TODO: enable once the attribute orgID gets added to the application model
	// t.Run("error_app_not_found_wrong_org_id", func(t *testing.T) {
	// 	appFixture := ff.NewApplication().WithEnvironmentID(env.ID).Save()
	// 	app, err := applicationRepo.FindByID(ctx, domain.ApplicationID(appFixture.ID), iam.NewOrgID())
	// 	require.ErrorIs(t, err, domain.ErrAppNotFound)
	// 	require.Nil(t, app)
	// })
}
