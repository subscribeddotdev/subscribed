package components_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed/server/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed/server/tests/client"
	"github.com/subscribeddotdev/subscribed/server/tests/fixture"
)

type JSON map[string]interface{}

func (j JSON) Marshall(t *testing.T) []byte {
	data, err := json.Marshal(j)
	require.NoError(t, err)
	return data
}

func TestMessages_SendMessage(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()
	app := ff.NewApplication().WithEnvironmentID(env.ID).Save()
	apiKey := ff.NewApiKey().WithOrgID(org.ID).WithEnvironmentID(env.ID).Save()
	eventType := ff.NewEventType().WithOrgID(org.ID).Save()
	ff.NewEndpoint().WithEventTypeIDs([]string{eventType.ID}).WithAppID(app.ID).Save()
	apiClient := getClientWithApiKey(t, apiKey.SecretKey)

	t.Run("send_message", func(t *testing.T) {
		payload := JSON{
			"id":         gofakeit.UUID(),
			"first_name": gofakeit.FirstName(),
			"last_name":  gofakeit.LastName(),
			"email":      gofakeit.Email(),
		}.Marshall(t)

		resp, err := apiClient.SendMessage(ctx, app.ID, client.SendMessageRequest{
			EventTypeId: eventType.ID,
			Payload:     string(payload),
		})
		require.NoError(t, err)
		require.Equal(t, http.StatusCreated, resp.StatusCode)
		msg, err := models.Messages(
			models.MessageWhere.ApplicationID.EQ(app.ID),
			models.MessageWhere.EventTypeID.EQ(eventType.ID),
		).One(ctx, db)
		require.NoError(t, err)

		assert.Equal(t, string(payload), msg.Payload)
	})
}
