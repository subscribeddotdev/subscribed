package command

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

type CreateApplication struct {
	Name  string
	OrgID iam.OrgID
	EnvID domain.EnvironmentID
}

type CreateApplicationHandler struct {
	repo domain.ApplicationRepository
}

func NewCreateApplicationHandler(repo domain.ApplicationRepository) CreateApplicationHandler {
	return CreateApplicationHandler{
		repo: repo,
	}
}

func (c CreateApplicationHandler) Execute(ctx context.Context, cmd CreateApplication) error {
	application, err := domain.NewApplication(cmd.Name, cmd.EnvID)
	if err != nil {
		return err
	}

	err = c.repo.Insert(ctx, application)
	if err != nil {
		return err
	}

	return nil
}
