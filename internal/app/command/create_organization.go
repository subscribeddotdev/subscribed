package command

import (
	"context"
	"fmt"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

type CreateOrganization struct {
	FirstName       string
	LastName        string
	Email           iam.Email
	LoginProviderID iam.LoginProviderID
}

func (o CreateOrganization) validate() error {
	if err := o.LoginProviderID.Validate(); err != nil {
		return err
	}

	if o.Email.IsEmpty() {
		return fmt.Errorf("email is empty")
	}

	return nil
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
	if err := cmd.validate(); err != nil {
		return err
	}

	return c.txProvider.Transact(ctx, func(adapters TransactableAdapters) error {
		exists, err := adapters.MemberRepository.ExistsByOr(ctx, cmd.Email, cmd.LoginProviderID)
		if err != nil {
			return err
		}

		if exists {
			// idempotent-friendliness
			return nil
		}

		org := iam.NewOrganization()

		err = adapters.OrganizationRepository.Insert(ctx, org)
		if err != nil {
			return fmt.Errorf("unable to save organization: %v", err)
		}

		defaultEnvironments, err := getDefaultEnvironments(org.ID())
		if err != nil {
			return fmt.Errorf("unable to generate the default environments: %v", err)
		}

		for _, env := range defaultEnvironments {
			err = adapters.EnvironmentRepository.Insert(ctx, env)
			if err != nil {
				return fmt.Errorf("unable to save the env '%s' due to the error: %v", env.Name(), err)
			}
		}

		member, err := iam.NewMember(org.ID(), cmd.LoginProviderID, cmd.FirstName, cmd.LastName, cmd.Email)
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

func getDefaultEnvironments(orgID iam.OrgID) ([]*domain.Environment, error) {
	prod, err := domain.NewEnvironment("Production", orgID.String(), domain.EnvTypeProduction)
	if err != nil {
		return nil, err
	}

	dev, err := domain.NewEnvironment("Development", orgID.String(), domain.EnvTypeDevelopment)
	if err != nil {
		return nil, err
	}

	return []*domain.Environment{prod, dev}, nil
}
