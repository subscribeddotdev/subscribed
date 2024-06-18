package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/subscribeddotdev/subscribed-backend/internal/app"
)

type handlers struct {
	application                  *app.App
	loginProviderWebhookVerifier LoginProviderWebhookVerifier
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
