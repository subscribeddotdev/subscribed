package command

import (
	"context"

	"github.com/subscribeddotdev/subscribed/server/internal/domain"
)

type CreateApplication struct {
	Name  string
	EnvID domain.EnvironmentID
}

type CreateApplicationHandler struct {
	repo domain.ApplicationRepository
}

func NewCreateApplicationHandler(
	repo domain.ApplicationRepository,
) CreateApplicationHandler {
	return CreateApplicationHandler{
		repo: repo,
	}
}

func (c CreateApplicationHandler) Execute(ctx context.Context, cmd CreateApplication) (domain.ApplicationID, error) {
	application, err := domain.NewApplication(cmd.Name, cmd.EnvID)
	if err != nil {
		return "", err
	}

	err = c.repo.Insert(ctx, application)
	if err != nil {
		return "", err
	}

	return application.ID(), nil
}
