package http

import (
	"net/http"

	"github.com/subscribeddotdev/subscribed-backend/internal/app"
)

type handlers struct {
	application                  *app.App
	loginProviderWebhookVerifier LoginProviderWebhookVerifier
}

type LoginProviderWebhookVerifier interface {
	Verify(payload []byte, headers http.Header) error
}
