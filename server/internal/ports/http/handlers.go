package http

import (
	"net/http"

	"github.com/friendsofgo/errors"
	"github.com/labstack/echo/v4"
	"github.com/subscribeddotdev/subscribed/server/internal/app"
	"github.com/subscribeddotdev/subscribed/server/internal/app/command"
	"github.com/subscribeddotdev/subscribed/server/internal/app/query"
	"github.com/subscribeddotdev/subscribed/server/internal/domain"
	"github.com/subscribeddotdev/subscribed/server/internal/domain/iam"
)

type handlers struct {
	application *app.App
	jwtSecret   string
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
	var body CreateApplicationJSONRequestBody
	err := c.Bind(&body)
	if err != nil {
		return NewHandlerErrorWithStatus(err, "error-parsing-the-body", http.StatusBadRequest)
	}

	apiKey, err := h.resolveApiKeyFromCtx(c)
	if err != nil {
		return NewHandlerError(err, "error-retrieving-api-key")
	}

	id, err := h.application.Command.CreateApplication.Execute(c.Request().Context(), command.CreateApplication{
		Name:  body.Name,
		EnvID: apiKey.EnvID(),
	})
	if err != nil {
		return NewHandlerError(err, "unable-to-create-application")
	}

	payload := CreateApplicationPayload{
		Id: id.String(),
	}

	return c.JSON(http.StatusCreated, payload)
}

func (h handlers) SendMessage(c echo.Context, applicationID string) error {
	apiKey, err := h.resolveApiKeyFromCtx(c)
	if err != nil {
		return NewHandlerError(err, "error-retrieving-org-id")
	}

	var body SendMessageRequest
	err = c.Bind(&body)
	if err != nil {
		return NewHandlerErrorWithStatus(err, "error-parsing-the-body", http.StatusBadRequest)
	}

	err = h.application.Command.SendMessage.Execute(c.Request().Context(), command.SendMessage{
		OrgID:         apiKey.OrgID(),
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
	claims, err := h.resolveJwtClaimsFromCtx(c)
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
		OrgID:         iam.OrgID(claims.OrganizationID),
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

func (h handlers) CreateApiKey(c echo.Context) error {
	claims, err := h.resolveJwtClaimsFromCtx(c)
	if err != nil {
		return err
	}

	var body CreateApiKeyRequest
	err = c.Bind(&body)
	if err != nil {
		return NewHandlerErrorWithStatus(err, "error-parsing-the-body", http.StatusBadRequest)
	}

	ak, err := h.application.Command.CreateApiKey.Execute(c.Request().Context(), command.CreateApiKey{
		Name:          body.Name,
		ExpiresAt:     body.ExpiresAt,
		EnvironmentID: domain.EnvironmentID(body.EnvironmentId),
		OrgID:         claims.OrganizationID,
	})
	if err != nil {
		return NewHandlerError(err, "unable-to-create-api-key")
	}

	return c.JSON(http.StatusCreated, CreateApiKeyPayload{UnmaskedApiKey: ak.SecretKey().FullKey()})
}

func (h handlers) GetEnvironments(c echo.Context) error {
	claims, err := h.resolveJwtClaimsFromCtx(c)
	if err != nil {
		return err
	}

	envs, err := h.application.Query.AllEnvironments.Execute(c.Request().Context(), query.AllEnvironments{
		OrgID: claims.OrganizationID,
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
	claims, err := h.resolveJwtClaimsFromCtx(c)
	if err != nil {
		return err
	}

	apiKeys, err := h.application.Query.AllApiKeys.Execute(c.Request().Context(), query.AllApiKeys{
		OrgID:         claims.OrganizationID,
		EnvironmentID: params.EnvironmentId,
	})
	if err != nil {
		return NewHandlerError(err, "error-fetching-api-keys")
	}

	data := make([]ApiKey, len(apiKeys))
	for i, apiKey := range apiKeys {
		data[i] = ApiKey{
			Id:              apiKey.Id().String(),
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

func (h handlers) DestroyApiKey(c echo.Context, apiKeyId string) error {
	claims, err := h.resolveJwtClaimsFromCtx(c)
	if err != nil {
		return err
	}

	err = h.application.Command.DestroyApiKey.Execute(c.Request().Context(), command.DestroyApiKey{
		OrgID: claims.OrganizationID,
		ID:    domain.ApiKeyID(apiKeyId),
	})
	if errors.Is(err, domain.ErrApiKeyNotFound) {
		return NewHandlerErrorWithStatus(err, "api-key-not-found", http.StatusNotFound)
	}
	if err != nil {
		return NewHandlerError(err, "error-destroying-api-key")
	}

	return c.NoContent(http.StatusNoContent)
}

func (h handlers) GetApplications(c echo.Context, params GetApplicationsParams) error {
	// TODO: Add support for resolving important details from when the request is made with an api key instead
	claims, err := h.resolveJwtClaimsFromCtx(c)
	if err != nil {
		return err
	}

	result, err := h.application.Query.AllApplications.Execute(c.Request().Context(), query.AllApplications{
		PaginationParams: query.NewPaginationParams(params.Page, params.Limit),
		EnvironmentID:    params.EnvironmentID,
		OrgID:            claims.OrganizationID,
	})
	if err != nil {
		return NewHandlerError(err, "error-retrieving-applications")
	}

	data := make([]Application, len(result.Data))

	for i, app := range result.Data {
		data[i] = Application{
			Name:          app.Name(),
			Id:            app.ID().String(),
			EnvironmentId: app.EnvID().String(),
			CreatedAt:     app.CreatedAt(),
		}
	}

	return c.JSON(http.StatusOK, GetApplicationsPayload{
		Data:       data,
		Pagination: mapToPaginationResponse(result),
	})
}

func (h handlers) GetApplicationById(c echo.Context, applicationID ApplicationId) error {
	claims, err := h.resolveJwtClaimsFromCtx(c)
	if err != nil {
		return err
	}

	app, err := h.application.Query.Application.Execute(c.Request().Context(), query.Application{
		ApplicationID: applicationID,
		OrgID:         claims.OrganizationID,
	})
	if errors.Is(err, domain.ErrAppNotFound) {
		return NewHandlerErrorWithStatus(err, "error-application-not-found", http.StatusNotFound)
	}

	if err != nil {
		return NewHandlerError(err, "error-retrieving-application")
	}

	return c.JSON(http.StatusOK, GetApplicationByIdPayload{Data: Application{
		Id:            app.ID().String(),
		Name:          app.Name(),
		EnvironmentId: app.EnvID().String(),
		CreatedAt:     app.CreatedAt(),
	}})
}

func (h handlers) GetEventTypes(c echo.Context, params GetEventTypesParams) error {
	claims, err := h.resolveJwtClaimsFromCtx(c)
	if err != nil {
		return err
	}

	result, err := h.application.Query.AllEventTypes.Execute(c.Request().Context(), query.AllEventTypes{
		PaginationParams: query.NewPaginationParams(params.Page, params.Limit),
		OrgID:            claims.OrganizationID,
	})
	if err != nil {
		return NewHandlerError(err, "error-retrieving-event-types")
	}

	data := make([]EventType, len(result.Data))

	for i, et := range result.Data {
		data[i] = EventType{
			Id:            et.ID().String(),
			Name:          et.Name(),
			Description:   et.Description(),
			Schema:        et.Schema(),
			SchemaExample: et.SchemaExample(),
			CreatedAt:     et.CreatedAt(),
			ArchivedAt:    et.ArchivedAt(),
		}
	}

	return c.JSON(http.StatusOK, GetEventTypesPayload{
		Data:       data,
		Pagination: mapToPaginationResponse(result),
	})
}

func (h handlers) resolveJwtClaimsFromCtx(c echo.Context) (*jwtCustomClaims, error) {
	claims, ok := c.Get("user_claims").(*jwtCustomClaims)
	if !ok {
		return nil, NewHandlerErrorWithStatus(errors.New("unable-to-retrieve-claims"), "unable-to-retrieve-claims", http.StatusUnauthorized)
	}

	return claims, nil
}

func (h handlers) resolveApiKeyFromCtx(c echo.Context) (*domain.ApiKey, error) {
	val := c.Get("api_key")
	if val == nil {
		return nil, errors.New("api_key hasn't been set in the context")
	}

	apiKey, ok := val.(*domain.ApiKey)
	if !ok {
		return nil, errors.New("invalid api_key type")
	}

	return apiKey, nil
}

func mapToPaginationResponse[D any](p query.Paginated[D]) Pagination {
	return Pagination{
		CurrentPage: p.CurrentPage,
		PerPage:     p.PerPage,
		Total:       p.Total,
		TotalPages:  p.TotalPages,
	}
}
