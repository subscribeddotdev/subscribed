package command

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type AddEndpoint struct {
	ApplicationID domain.ApplicationID
	EndpointUrl   string
	Description   string
	EventTypeIDs  []domain.EventTypeID
}

type AddEndpointHandler struct {
	repo domain.EndpointRepository
}

func NewAddEndpointHandler(repo domain.EndpointRepository) AddEndpointHandler {
	return AddEndpointHandler{
		repo: repo,
	}
}

func (c AddEndpointHandler) Execute(ctx context.Context, cmd AddEndpoint) error {
	endpointURL, err := domain.NewEndpointURL(cmd.EndpointUrl)
	if err != nil {
		return err
	}

	endpoint, err := domain.NewEndpoint(endpointURL, cmd.ApplicationID, cmd.Description, cmd.EventTypeIDs)
	if err != nil {
		return err
	}

	return c.repo.Insert(ctx, endpoint)
}
