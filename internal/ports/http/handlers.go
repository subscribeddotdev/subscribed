package http

import (
	"net/http"

	"github.com/subscribeddotdev/subscribed-backend/internal/app"
)

type handlers struct {
	application                  *app.App
	loginProviderWebhookVerifier loginProviderWebhookVerifier
}

type loginProviderWebhookVerifier interface {
	Verify(payload []byte, headers http.Header) error
}
