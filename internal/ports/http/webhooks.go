package http

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/subscribeddotdev/subscribed-backend/internal/app/command"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

func (h handlers) CreateAccount(c echo.Context, params CreateAccountParams) error {
	rawBody, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return NewHandlerError(err, "error-reading-the-body")
	}
	err = c.Request().Body.Close()
	if err != nil {
		return NewHandlerError(err, "error-closing-the-body")
	}

	err = h.loginProviderWebhookVerifier.Verify(rawBody, c.Request().Header)
	if err != nil {
		return NewHandlerErrorWithStatus(err, "webhook-auth-failed", http.StatusForbidden)
	}

	var body CreateAccountRequest
	err = json.Unmarshal(rawBody, &body)
	if err != nil {
		return NewHandlerError(err, "error-binding-the-body")
	}

	firstName := ""
	lastName := ""

	if body.Data.FirstName != nil {
		firstName = *body.Data.FirstName
	}

	if body.Data.LastName != nil {
		lastName = *body.Data.LastName
	}

	if len(body.Data.EmailAddresses) == 0 {
		return NewHandlerErrorWithStatus(err, "missing-email-address", http.StatusBadRequest)
	}

	email, err := iam.NewEmail(body.Data.EmailAddresses[0].EmailAddress)
	if err != nil {
		return NewHandlerErrorWithStatus(err, "malformed-email-address", http.StatusBadRequest)
	}

	err = h.application.Command.CreateOrganization.Execute(c.Request().Context(), command.CreateOrganization{
		FirstName:       firstName,
		LastName:        lastName,
		Email:           email,
		LoginProviderID: iam.LoginProviderID(body.Data.Id),
	})
	if err != nil {
		return NewHandlerError(err, "unable-to-create-account")
	}

	return c.NoContent(http.StatusCreated)
}
