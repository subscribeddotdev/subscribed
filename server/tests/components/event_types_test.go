package components_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed/server/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed/server/tests"
	"github.com/subscribeddotdev/subscribed/server/tests/client"
	"github.com/subscribeddotdev/subscribed/server/tests/fixture"
)

func TestEventTypes(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	member, password := ff.NewMember().WithOrganizationID(org.ID).Save()
	token := signIn(t, member.Email, password)
	apiClient := getClientWithToken(t, token)

	t.Run("create_event_type", func(t *testing.T) {
		resp, err := apiClient.CreateEventType(ctx, client.CreateEventTypeRequest{
			Name:        fmt.Sprintf("%s.%s", gofakeit.Noun(), gofakeit.Verb()),
			Description: toPtr(gofakeit.Sentence(20)),
		})
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
	})

	t.Run("get_all_event_types_by_org_id", func(t *testing.T) {
		eventTypesFixtureCount := 20
		eventTypesFixture := make([]models.EventType, eventTypesFixtureCount)

		for i := 0; i < eventTypesFixtureCount; i++ {
			eventTypesFixture[i] = ff.NewEventType().WithOrgID(org.ID).Save()
		}

		perPage := 5
		totalPages := eventTypesFixtureCount / 5

		for i := 0; i < totalPages; i++ {
			currentPage := i + 1
			resp, err := apiClient.GetEventTypesWithResponse(ctx, &client.GetEventTypesParams{
				Limit: &perPage,
				Page:  &currentPage,
			})

			require.NoError(t, err)
			require.Equal(t, http.StatusOK, resp.StatusCode())
			require.NotEmpty(t, resp.JSON200.Data)
			require.Equal(t, currentPage, resp.JSON200.Pagination.CurrentPage)
			require.GreaterOrEqual(t, resp.JSON200.Pagination.Total, eventTypesFixtureCount)
			require.Equal(t, perPage, resp.JSON200.Pagination.PerPage)
			require.Equal(t, totalPages, resp.JSON200.Pagination.TotalPages)
		}
	})

	t.Run("get_event_type_by_id", func(t *testing.T) {
		et := ff.NewEventType().WithArchivedAt(gofakeit.PastDate()).WithOrgID(org.ID).Save()

		resp, err := apiClient.GetEventTypeByIdWithResponse(ctx, et.ID)
		require.NoError(t, err)
		require.Equal(t, http.StatusOK, resp.StatusCode())
		require.Equal(t, et.ID, resp.JSON200.Data.Id)
		require.Equal(t, et.Name, resp.JSON200.Data.Name)
		require.Equal(t, et.Description.String, resp.JSON200.Data.Description)
		require.Equal(t, et.Schema.String, resp.JSON200.Data.Schema)
		require.Equal(t, et.SchemaExample.String, resp.JSON200.Data.SchemaExample)
		tests.RequireEqualTime(t, et.CreatedAt, resp.JSON200.Data.CreatedAt)
		tests.RequireEqualTime(t, et.ArchivedAt.Time, *resp.JSON200.Data.ArchivedAt)

		resp, err = apiClient.GetEventTypeByIdWithResponse(ctx, "non-existent-id")
		require.NoError(t, err)
		require.Equal(t, http.StatusNotFound, resp.StatusCode())
	})
}
