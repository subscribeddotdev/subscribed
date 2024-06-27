package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/subscribeddotdev/subscribed-backend/internal/app"
	"github.com/subscribeddotdev/subscribed-backend/internal/app/command"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

type handlers struct {
	application                  *app.App
	loginProviderWebhookVerifier LoginProviderWebhookVerifier
}

func (h handlers) AddEndpoint(c echo.Context, applicationID string) error {
	appID, err := domain.NewIdFromString(applicationID)
	if err != nil {
		return NewHandlerErrorWithStatus(err, "invalid-application-id", http.StatusBadRequest)
	}

	var body AddEndpointRequest
	err = c.Bind(&body)
	if err != nil {
		return NewHandlerErrorWithStatus(err, "error-parsing-the-body", http.StatusBadRequest)
	}

	var description string
	if body.Description != nil {
		description = *body.Description
	}

	err = h.application.Command.AddEndpoint.Execute(c.Request().Context(), command.AddEndpoint{
		ApplicationID:          appID,
		EndpointUrl:            body.Url,
		Description:            description,
		EventTypesSubscribedTo: nil,
	})
	if err != nil {
		return NewHandlerError(err, "unable-to-add-endpoint")
	}

	return c.NoContent(http.StatusNoContent)
}

type LoginProviderWebhookVerifier interface {
	Verify(payload []byte, headers http.Header) error
}

func (h handlers) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

func (h handlers) CreateApplication(c echo.Context) error {
	return c.NoContent(http.StatusCreated)
}

func (h handlers) SendMessage(ctx echo.Context, applicationID string) error {
	//TODO implement me
	panic("implement me")
}
