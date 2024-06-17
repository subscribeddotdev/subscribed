package command

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type CreateApplication struct {
	Name  string
	OrgID domain.ID
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
	application, err := domain.NewApplication(cmd.Name, cmd.OrgID)
	if err != nil {
		return err
	}

	err = c.repo.Insert(ctx, application)
	if err != nil {
		return err
	}

	return nil
}
