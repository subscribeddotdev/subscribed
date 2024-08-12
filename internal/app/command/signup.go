package command

import (
	"context"
	"fmt"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

type Signup struct {
	FirstName string
	LastName  string
	Email     iam.Email
	Password  iam.Password
}

type SignupHandler struct {
	txProvider TransactionProvider
}

func NewSignupHandler(txProvider TransactionProvider) SignupHandler {
	return SignupHandler{
		txProvider: txProvider,
	}
}

func (c SignupHandler) Execute(ctx context.Context, cmd Signup) error {
	return c.txProvider.Transact(ctx, func(adapters TransactableAdapters) error {
		org := iam.NewOrganization()

		err := adapters.OrganizationRepository.Insert(ctx, org)
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

		member, err := iam.NewMember(org.ID(), cmd.FirstName, cmd.LastName, cmd.Email, cmd.Password)
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
