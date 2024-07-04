package transaction

import (
	"context"
	"fmt"

	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/psql"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/events"
	"github.com/subscribeddotdev/subscribed-backend/internal/app/command"
	"github.com/subscribeddotdev/subscribed-backend/internal/common/logs"
)

type PsqlProvider struct {
	logger         *logs.Logger
	db             boil.ContextBeginner
	eventPublisher *events.Publisher
}

func NewPsqlProvider(db boil.ContextBeginner, eventPublisher *events.Publisher, logger *logs.Logger) PsqlProvider {
	return PsqlProvider{
		logger:         logger,
		db:             db,
		eventPublisher: eventPublisher,
	}
}

func (p PsqlProvider) Transact(ctx context.Context, f command.TransactFunc) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("unable to begin transaction: %v", err)
	}

	adapters := command.TransactableAdapters{
		EventPublisher:         p.eventPublisher,
		MemberRepository:       psql.NewMemberRepository(tx),
		OrganizationRepository: psql.NewOrganizationRepository(tx),
		EnvironmentRepository:  psql.NewEnvironmentRepository(tx),
		MessageRepository:      psql.NewMessageRepository(tx),
		EndpointRepository:     psql.NewEndpointRepository(tx),
	}

	if err = f(adapters); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			p.logger.Error("Rollback error", "rollback_err", rollbackErr)
		}

		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
