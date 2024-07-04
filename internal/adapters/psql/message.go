package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/friendsofgo/errors"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
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
		OrgID:         message.OrgID().String(),
		ApplicationID: message.ApplicationID().String(),
		EventTypeID:   message.EventTypeID().String(),
		Payload:       message.Payload(),
		SentAt:        message.SentAt(),
	}

	return model.Insert(ctx, o.db, boil.Infer())
}

func (o MessageRepository) ByID(ctx context.Context, id domain.ID) (*domain.Message, error) {
	model, err := models.FindMessage(ctx, o.db, id.String())
	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrMessageNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("error querying message by id '%s': %v", id, err)
	}

	return domain.UnMarshallMessage(model.ID, model.EventTypeID, model.ApplicationID, model.OrgID, model.SentAt, model.Payload)
}
