package psql

import (
	"context"

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
	return nil
}
