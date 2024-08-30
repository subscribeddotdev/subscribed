package command

import (
	"context"
	"fmt"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type CreateApplication struct {
	Name     string
	ApiKeyID domain.ApiKeyID
}

type CreateApplicationHandler struct {
	repo         domain.ApplicationRepository
	apiKeyFinder apiKeyFinder
}

type apiKeyFinder interface {
	FindByID(ctx context.Context, id domain.ApiKeyID) (*domain.ApiKey, error)
}

func NewCreateApplicationHandler(
	repo domain.ApplicationRepository,
	apiKeyFinder apiKeyFinder,
) CreateApplicationHandler {
	return CreateApplicationHandler{
		repo:         repo,
		apiKeyFinder: apiKeyFinder,
	}
}

func (c CreateApplicationHandler) Execute(ctx context.Context, cmd CreateApplication) (domain.ApplicationID, error) {
	apiKey, err := c.apiKeyFinder.FindByID(ctx, cmd.ApiKeyID)
	if err != nil {
		return "", fmt.Errorf("unable to find the api key: %v", err)
	}

	application, err := domain.NewApplication(cmd.Name, apiKey.EnvID())
	if err != nil {
		return "", err
	}

	err = c.repo.Insert(ctx, application)
	if err != nil {
		return "", err
	}

	return application.ID(), nil
}
