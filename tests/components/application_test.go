package components_test

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/tests/client"
)

func TestApplicationLifecycle(t *testing.T) {
	t.Run("send_message", func(t *testing.T) {
		payload, err := gofakeit.JSON(&gofakeit.JSONOptions{
			Type: "object",
			Fields: []gofakeit.Field{
				{Name: "first_name", Function: "firstname"},
				{Name: "last_name", Function: "lastname"},
			},
		})
		require.NoError(t, err)

		appID := domain.NewID().String()
		eventTypeID := domain.NewID().String()
		resp, err := getClient(t, "").SendMessage(ctx, appID, client.SendMessageRequest{
			EventTypeId: eventTypeID, //TODO: replace this with an existing eventTypeID
			Payload:     string(payload),
		})
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp.StatusCode)
		msg, err := models.Messages(
			models.MessageWhere.ApplicationID.EQ(appID),
			models.MessageWhere.EventTypeID.EQ(eventTypeID),
		).One(ctx, db)
		require.NoError(t, err)

		assert.Equal(t, string(payload), msg.Payload)
	})
}
