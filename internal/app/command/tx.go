package command

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

type TransactFunc func(adapters TransactableAdapters) error

type TransactionProvider interface {
	Transact(ctx context.Context, f TransactFunc) error
}

type TransactableAdapters struct {
	EventPublisher         EventPublisher
	MemberRepository       iam.MemberRepository
	OrganizationRepository iam.OrganizationRepository
	EnvironmentRepository  domain.EnvironmentRepository
	MessageRepository      domain.MessageRepository
	EndpointRepository     domain.EndpointRepository
}

type EventPublisher interface {
	PublishMessageSent(ctx context.Context, e MessageSent) error
}
