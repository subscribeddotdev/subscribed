package command

import (
	"context"

	"github.com/subscribeddotdev/subscribed/server/internal/domain"
)

type ArchiveEventType struct {
	OrgID       string
	EventTypeID domain.EventTypeID
}

type ArchiveEventTypeHandler struct {
	repo domain.EventTypeRepository
}

func NewArchiveEventTypeHandler(repo domain.EventTypeRepository) ArchiveEventTypeHandler {
	return ArchiveEventTypeHandler{
		repo: repo,
	}
}

func (h ArchiveEventTypeHandler) Execute(ctx context.Context, cmd ArchiveEventType) error {
	return h.repo.Update(ctx, cmd.OrgID, cmd.EventTypeID, func(e *domain.EventType) error {
		e.Archive()
		return nil
	})
}
