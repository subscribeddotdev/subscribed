package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/subscribeddotdev/subscribed-backend/internal/app/command"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

func (h handlers) SignUp(c echo.Context) error {
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

	err = h.application.Command.SignUp.Execute(c.Request().Context(), command.Signup{
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

func (h handlers) SignIn(c echo.Context) error {
	var body SigninRequest
	err := c.Bind(&body)
	if err != nil {
		return NewHandlerError(err, "error-binding-the-body")
	}

	email, err := iam.NewEmail(body.Email)
	if err != nil {
		return NewHandlerErrorWithStatus(err, "malformed-email-address", http.StatusBadRequest)
	}

	member, err := h.application.Command.SignIn.Execute(c.Request().Context(), command.SignIn{
		Email:             email,
		PlainTextPassword: body.Password,
	})
	if err != nil {
		return NewHandlerError(err, "error-authenticating")
	}

	token, err := h.jwtIssuer.Issue(member)
	if err != nil {
		return NewHandlerError(err, "error-signing-jwt")
	}

	return c.JSON(http.StatusOK, SignInPayload{token})
}
