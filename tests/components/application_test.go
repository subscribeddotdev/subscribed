package components_test

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/tests/client"
	"github.com/subscribeddotdev/subscribed-backend/tests/fixture"
)

func TestApplication_Lifecycle(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()
	app := ff.NewApplication().WithEnvironmentID(env.ID).Save()
	eventType := ff.NewEventType().WithOrgID(org.ID).Save()

	t.Run("send_message", func(t *testing.T) {
		payload, err := gofakeit.JSON(&gofakeit.JSONOptions{
			Type: "object",
			Fields: []gofakeit.Field{
				{Name: "first_name", Function: "firstname"},
				{Name: "last_name", Function: "lastname"},
			},
		})
		require.NoError(t, err)

		resp, err := getClient(t, "").SendMessage(ctx, app.ID, client.SendMessageRequest{
			EventTypeId: eventType.ID,
			Payload:     string(payload),
		})
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		msg, err := models.Messages(
			models.MessageWhere.ApplicationID.EQ(app.ID),
			models.MessageWhere.EventTypeID.EQ(eventType.ID),
		).One(ctx, db)
		require.NoError(t, err)

		assert.Equal(t, string(payload), msg.Payload)
	})
}
