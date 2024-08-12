package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/subscribeddotdev/subscribed-backend/internal/app/command"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

func (h handlers) Signup(c echo.Context) error {
	var body SignupRequest
	err := c.Bind(&body)
	if err != nil {
		return NewHandlerError(err, "error-binding-the-body")
	}

	email, err := iam.NewEmail(body.Email)
	if err != nil {
		return NewHandlerErrorWithStatus(err, "malformed-email-address", http.StatusBadRequest)
	}

	password, err := iam.NewPassword(body.Password)
	if err != nil {
		return NewHandlerErrorWithStatus(err, "malformed-password", http.StatusBadRequest)
	}

	err = h.application.Command.CreateOrganization.Execute(c.Request().Context(), command.Signup{
		Email:     email,
		Password:  password,
		FirstName: body.FirstName,
		LastName:  body.LastName,
	})
	if err != nil {
		return NewHandlerError(err, "unable-to-create-account")
	}

	return c.NoContent(http.StatusCreated)
}
