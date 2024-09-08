package components_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed/server/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed/server/internal/domain"
	"github.com/subscribeddotdev/subscribed/server/tests"
	"github.com/subscribeddotdev/subscribed/server/tests/client"
	"github.com/subscribeddotdev/subscribed/server/tests/fixture"
)

func TestApplications(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	member, password := ff.NewMember().WithOrganizationID(org.ID).Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()
	token := signIn(t, member.Email, password)
	apiClient := getClientWithToken(t, token)
	applicationFixtureCount := 20
	appsFixture := make([]models.Application, applicationFixtureCount)

	for i := 0; i < applicationFixtureCount; i++ {
		appsFixture[i] = ff.NewApplication().WithEnvironmentID(env.ID).Save()
	}

	t.Run("get_all_applications_by_env_id", func(t *testing.T) {
		perPage := 5
		totalPages := applicationFixtureCount / 5

		for i := 0; i < totalPages; i++ {
			currentPage := i + 1
			resp, err := apiClient.GetApplicationsWithResponse(ctx, &client.GetApplicationsParams{
				EnvironmentID: env.ID,
				Limit:         &perPage,
				Page:          &currentPage,
			})

			require.NoError(t, err)
			require.Equal(t, http.StatusOK, resp.StatusCode())
			require.NotEmpty(t, resp.JSON200.Data)
			require.Equal(t, currentPage, resp.JSON200.Pagination.CurrentPage)
			require.Equal(t, applicationFixtureCount, resp.JSON200.Pagination.Total)
			require.Equal(t, perPage, resp.JSON200.Pagination.PerPage)
			require.Equal(t, totalPages, resp.JSON200.Pagination.TotalPages)
		}
	})

	t.Run("get_application_by_id", func(t *testing.T) {
		appFixture := appsFixture[0]
		resp, err := apiClient.GetApplicationByIdWithResponse(ctx, appFixture.ID)

		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode())
		require.Equal(t, appFixture.ID, resp.JSON200.Data.Id)
		require.Equal(t, appFixture.Name, resp.JSON200.Data.Name)
		require.Equal(t, appFixture.EnvironmentID, resp.JSON200.Data.EnvironmentId)
		tests.RequireEqualTime(t, appFixture.CreatedAt, resp.JSON200.Data.CreatedAt)

		// Not found
		resp, err = apiClient.GetApplicationByIdWithResponse(ctx, domain.NewApplicationID().String())
		require.NoError(t, err)
		require.Equal(t, http.StatusNotFound, resp.StatusCode())
	})
}
