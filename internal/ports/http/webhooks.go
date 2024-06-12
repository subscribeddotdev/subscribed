package http

import (
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h handlers) CreateAccount(c echo.Context, params CreateAccountParams) error {
	rawBody, err := io.ReadAll(c.Request().Body)
	if err != nil {
		return NewHandlerError(err, "error-reading-the-body")
	}
	// hello
	err = h.loginProviderWebhookVerifier.Verify(rawBody, c.Request().Header)
	if err != nil {
		return NewHandlerErrorWithStatus(err, "webhook-auth-failed", http.StatusForbidden)
	}

	return c.NoContent(http.StatusCreated)
}
