package components_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/tests/client"
	"github.com/subscribeddotdev/subscribed-backend/tests/fixture"
)

func TestApplications(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	member, password := ff.NewMember().WithOrganizationID(org.ID).Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()
	token := signIn(t, member.Email, password)
	apiClient := getClientWithToken(t, token)
	applicationFixtureCount := 20
	for i := 0; i < applicationFixtureCount; i++ {
		ff.NewApplication().WithEnvironmentID(env.ID).Save()
	}

	t.Run("get_all_applications_by_env_id", func(t *testing.T) {
		perPage := 5
		totalPages := applicationFixtureCount / 5

		for i := 0; i < totalPages; i++ {
			currentPage := i + 1
			resp, err := apiClient.GetApplicationsWithResponse(ctx, &client.GetApplicationsParams{
				EnvironmentId: env.ID,
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
}
