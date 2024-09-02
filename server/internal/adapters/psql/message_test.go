package psql_test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/tests/fixture"
)

func TestMessageRepository_Lifecycle(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	eventType := ff.NewEventType().WithOrgID(org.ID).Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()
	app := ff.NewApplication().WithEnvironmentID(env.ID).Save()

	msg, err := domain.NewMessage(
		domain.EventTypeID(eventType.ID),
		org.ID,
		domain.ApplicationID(app.ID),
		gofakeit.Sentence(10),
	)
	require.NoError(t, err)

	t.Run("insert_msg", func(t *testing.T) {
		require.NoError(t, msgRepo.Insert(ctx, msg))
	})

	t.Run("find_by_id", func(t *testing.T) {
		foundMsg, err := msgRepo.ByID(ctx, msg.Id())
		require.NoError(t, err)

		assert.Equal(t, msg.EventTypeID().String(), foundMsg.EventTypeID().String())
		assert.Equal(t, msg.Payload(), foundMsg.Payload())
		assert.Equal(t, msg.ApplicationID().String(), foundMsg.ApplicationID().String())
		assert.Equal(t, msg.SentAt().UTC().Truncate(time.Second), foundMsg.SentAt().Truncate(time.Second))
	})

	t.Run("error_not_found_by_id", func(t *testing.T) {
		foundMsg, err := msgRepo.ByID(ctx, domain.NewMessageID())
		require.Nil(t, foundMsg)
		assert.ErrorIs(t, err, domain.ErrMessageNotFound)
	})

	t.Run("save_message_send_attempt", func(t *testing.T) {
		// TODO: implement me
	})
}
