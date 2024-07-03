package psql

import (
	"context"

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
