package psql

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/friendsofgo/errors"
	"github.com/subscribeddotdev/subscribed/server/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed/server/internal/domain"
	"github.com/subscribeddotdev/subscribed/server/tests"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type MessageRepository struct {
	db boil.ContextExecutor
}

func NewMessageRepository(db boil.ContextExecutor) *MessageRepository {
	return &MessageRepository{
		db: db,
	}
}

func (o MessageRepository) Insert(ctx context.Context, message *domain.Message) error {
	model := models.Message{
		ID:            message.Id().String(),
		OrgID:         message.OrgID(),
		ApplicationID: message.ApplicationID().String(),
		EventTypeID:   message.EventTypeID().String(),
		Payload:       message.Payload(),
		SentAt:        message.SentAt(),
	}

	return model.Insert(ctx, o.db, boil.Infer())
}

func (o MessageRepository) ByID(ctx context.Context, id domain.MessageID) (*domain.Message, error) {
	model, err := models.FindMessage(ctx, o.db, id.String())
	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrMessageNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("error querying message by id '%s': %v", id, err)
	}

	return domain.UnMarshallMessage(
		domain.MessageID(model.ID),
		tests.ToPtr(domain.EventTypeID(model.EventTypeID)),
		domain.ApplicationID(model.ApplicationID),
		model.OrgID,
		model.SentAt,
		model.Payload,
	)
}

func (o MessageRepository) SaveMessageSendAttempt(ctx context.Context, attempt *domain.MessageSendAttempt) error {
	var err error
	var reqHeaders []byte

	if attempt.RequestHeaders() != nil {
		reqHeaders, err = json.Marshal(attempt.RequestHeaders())
		if err != nil {
			return err
		}
	}

	model := models.MessageSendAttempt{
		ID:             attempt.ID().String(),
		MessageID:      attempt.MessageID().String(),
		EndpointID:     attempt.EndpointID().String(),
		AttemptedAt:    attempt.Timestamp(),
		Status:         attempt.Status().String(),
		Response:       null.StringFrom(attempt.Response()),
		StatusCode:     null.Int16From(int16(attempt.StatusCode())),
		RequestHeaders: null.JSONFrom(reqHeaders),
	}

	err = model.Insert(ctx, o.db, boil.Infer())
	if err != nil {
		return fmt.Errorf("error saving messageSendAttempt of msg id '%s' due to: %v", attempt.MessageID(), err)
	}

	return nil
}
