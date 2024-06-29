package http

import (
	"net/http"

	"github.com/friendsofgo/errors"
	"github.com/labstack/echo/v4"
	"github.com/subscribeddotdev/subscribed-backend/internal/app"
	"github.com/subscribeddotdev/subscribed-backend/internal/app/command"
	"github.com/subscribeddotdev/subscribed-backend/internal/common/clerkhttp"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
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

func (h handlers) SendMessage(c echo.Context, applicationID string) error {
	var body SendMessageRequest
	err := c.Bind(&body)
	if err != nil {
		return NewHandlerErrorWithStatus(err, "error-parsing-the-body", http.StatusBadRequest)
	}

	err = h.application.Command.SendMessage.Execute(c.Request().Context(), command.SendMessage{
		ApplicationID: applicationID,
		EventTypeID:   body.EventTypeId,
		Payload:       body.Payload,
	})
	if err != nil {
		return NewHandlerError(err, "unable-to-send-message")
	}

	return c.NoContent(http.StatusCreated)
}

func (h handlers) CreateEventType(c echo.Context) error {
	member, err := h.resolveMemberFromCtx(c)
	if err != nil {
		return err
	}

	var body CreateEventTypeRequest
	err = c.Bind(&body)
	if err != nil {
		return NewHandlerErrorWithStatus(err, "error-parsing-the-body", http.StatusBadRequest)
	}

	var schema string
	var description string
	var schemaExample string

	if body.Description != nil {
		description = *body.Description
	}

	if body.Schema != nil {
		schema = *body.Schema
	}

	if body.SchemaExample != nil {
		schemaExample = *body.SchemaExample
	}

	err = h.application.Command.CreateEventType.Execute(c.Request().Context(), command.CreateEventType{
		OrgID:         member.OrganizationID().String(),
		Name:          body.Name,
		Description:   description,
		Schema:        schema,
		SchemaExample: schemaExample,
	})
	if err != nil {
		return NewHandlerError(err, "unable-to-send-message")
	}

	return c.NoContent(http.StatusCreated)
}

func (h handlers) resolveMemberFromCtx(c echo.Context) (*iam.Member, error) {
	claims, ok := clerkhttp.SessionClaimsFromContext(c)
	if !ok {
		return nil, NewHandlerErrorWithStatus(errors.New("unauthorized"), "unauthorized", http.StatusUnauthorized)
	}

	m, err := h.application.Authorization.ResolveMemberByLoginProviderID(c.Request().Context(), claims.Subject)
	if errors.Is(err, iam.ErrMemberNotFound) {
		return nil, NewHandlerErrorWithStatus(errors.New("forbidden"), "member-not-found", http.StatusForbidden)
	}

	if err != nil {
		return nil, NewHandlerErrorWithStatus(errors.New("unauthorized"), "unauthorized", http.StatusForbidden)
	}

	return m, nil
}
