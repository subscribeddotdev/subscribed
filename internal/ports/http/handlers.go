package http

import (
	"net/http"

	"github.com/friendsofgo/errors"
	"github.com/labstack/echo/v4"
	"github.com/subscribeddotdev/subscribed-backend/internal/app"
	"github.com/subscribeddotdev/subscribed-backend/internal/app/command"
	"github.com/subscribeddotdev/subscribed-backend/internal/app/query"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

type jwtIssuer interface {
	Issue(member *iam.Member) (string, error)
}

type handlers struct {
	application *app.App
	jwtIssuer   jwtIssuer
}

func (h handlers) HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "ok")
}

func (h handlers) AddEndpoint(c echo.Context, applicationID string) error {
	var body AddEndpointRequest
	err := c.Bind(&body)
	if err != nil {
		return NewHandlerErrorWithStatus(err, "error-parsing-the-body", http.StatusBadRequest)
	}

	var description string
	if body.Description != nil {
		description = *body.Description
	}

	var eventTypeIDs []domain.EventTypeID

	if body.EventTypeIds != nil {
		for _, eventTypeID := range *body.EventTypeIds {
			eID, err := domain.NewIdFromString(eventTypeID)
			if err != nil {
				return NewHandlerErrorWithStatus(err, "error-mapping-event-type-id", http.StatusBadRequest)
			}

			eventTypeIDs = append(eventTypeIDs, domain.EventTypeID(eID))
		}
	}

	err = h.application.Command.AddEndpoint.Execute(c.Request().Context(), command.AddEndpoint{
		ApplicationID: domain.ApplicationID(applicationID),
		EndpointUrl:   body.Url,
		Description:   description,
		EventTypeIDs:  eventTypeIDs,
	})
	if err != nil {
		return NewHandlerError(err, "unable-to-add-endpoint")
	}

	return c.NoContent(http.StatusNoContent)
}

func (h handlers) CreateApplication(c echo.Context) error {
	return c.NoContent(http.StatusCreated)
}

func (h handlers) SendMessage(c echo.Context, applicationID string) error {
	orgID, err := h.resolveOrgIdFromCtx(c)
	if err != nil {
		return NewHandlerError(err, "error-retrieving-org-id")
	}

	var body SendMessageRequest
	err = c.Bind(&body)
	if err != nil {
		return NewHandlerErrorWithStatus(err, "error-parsing-the-body", http.StatusBadRequest)
	}

	err = h.application.Command.SendMessage.Execute(c.Request().Context(), command.SendMessage{
		OrgID:         orgID,
		ApplicationID: domain.ApplicationID(applicationID),
		EventTypeID:   domain.EventTypeID(body.EventTypeId),
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
		OrgID:         member.OrgID(),
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

func (h handlers) CreateApiKey(c echo.Context, params CreateApiKeyParams) error {
	member, err := h.resolveMemberFromCtx(c)
	if err != nil {
		return err
	}

	var body CreateApiKeyRequest
	err = c.Bind(&body)
	if err != nil {
		return NewHandlerErrorWithStatus(err, "error-parsing-the-body", http.StatusBadRequest)
	}

	err = h.application.Command.CreateApiKey.Execute(c.Request().Context(), command.CreateApiKey{
		Name:          body.Name,
		ExpiresAt:     body.ExpiresAt,
		EnvironmentID: domain.EnvironmentID(params.EnvironmentId),
		OrgID:         member.OrgID().String(),
	})
	if err != nil {
		return NewHandlerError(err, "unable-to-create-api-key")
	}

	return c.NoContent(http.StatusCreated)
}

func (h handlers) GetEnvironments(c echo.Context) error {
	member, err := h.resolveMemberFromCtx(c)
	if err != nil {
		return err
	}

	envs, err := h.application.Query.Environments.Execute(c.Request().Context(), query.Environments{
		OrgID: member.OrgID().String(),
	})
	if err != nil {
		return err
	}

	data := make([]Environment, len(envs))

	for i, env := range envs {
		data[i] = Environment{
			ArchivedAt:     env.ArchivedAt(),
			CreatedAt:      env.CreatedAt(),
			Id:             env.ID().String(),
			Name:           env.Name(),
			OrganizationId: env.OrgID(),
			Type:           EnvironmentType(env.Type().String()),
		}
	}

	return c.JSON(http.StatusOK, GetAllEnvironmentsPayload{Data: data})
}

func (h handlers) GetAllApiKeys(c echo.Context, params GetAllApiKeysParams) error {
	member, err := h.resolveMemberFromCtx(c)
	if err != nil {
		return err
	}

	apiKeys, err := h.application.Query.AllApiKeys.Execute(c.Request().Context(), query.AllApiKeys{
		OrgID:         member.OrgID().String(),
		EnvironmentID: params.EnvironmentId,
	})
	if err != nil {
		return NewHandlerError(err, "error-fetching-api-keys")
	}

	data := make([]ApiKey, len(apiKeys))
	for i, apiKey := range apiKeys {
		data[i] = ApiKey{
			CreatedAt:       apiKey.CreatedAt(),
			EnvironmentId:   apiKey.EnvID().String(),
			ExpiresAt:       apiKey.ExpiresAt(),
			MaskedSecretKey: apiKey.SecretKey().String(),
			Name:            apiKey.Name(),
			OrganizationId:  apiKey.OrgID(),
		}
	}

	return c.JSON(http.StatusOK, GetAllApiKeysPayload{Data: data})
}

func (h handlers) resolveMemberFromCtx(c echo.Context) (*iam.Member, error) {
	panic("implement me")
	return &iam.Member{}, nil
}

func (h handlers) resolveOrgIdFromCtx(c echo.Context) (string, error) {
	val := c.Get("org_id")
	if val == nil {
		return "", errors.New("orgID hasn't been set in the context")
	}

	orgID, ok := val.(string)
	if !ok {
		return "", errors.New("invalid orgID type")
	}

	return orgID, nil
}
