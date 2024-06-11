package command

import (
	"context"
	"fmt"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

type CreateOrganization struct {
	FirstName       string
	LastName        string
	Email           iam.Email
	LoginProviderID iam.LoginProviderID
}

type CreateOrganizationHandler struct {
	txProvider TransactionProvider
}

func NewCreateOrganizationHandler(txProvider TransactionProvider) CreateOrganizationHandler {
	return CreateOrganizationHandler{
		txProvider: txProvider,
	}
}

func (c CreateOrganizationHandler) Execute(ctx context.Context, cmd CreateOrganization) error {
	return c.txProvider.Transact(ctx, func(adapters TransactableAdapters) error {
		org := iam.NewOrganization()

		err := adapters.OrganizationRepository.Insert(ctx, org)
		if err != nil {
			return fmt.Errorf("unable to save organization: %v", err)
		}

		member, err := iam.NewMember(org.Id(), cmd.LoginProviderID, cmd.FirstName, cmd.LastName, cmd.Email)
		if err != nil {
			return err
		}

		err = adapters.MemberRepository.Insert(ctx, member)
		if err != nil {
			return fmt.Errorf("unable to save member: %v", err)
		}

		return nil
	})
}
