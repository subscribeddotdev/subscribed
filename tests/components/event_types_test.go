package components_test

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/tests/client"
	"github.com/subscribeddotdev/subscribed-backend/tests/jwks"
)

func TestEventTypes(t *testing.T) {
	token := jwks.JwtGenerator(t, map[string]any{
		"sid": gofakeit.UUID(),
		"sub": "user_123",
		"iss": "https://clerk.com",
	})

	resp, err := getClient(t, token).CreateEventType(ctx, client.CreateEventTypeRequest{
		Name:        gofakeit.AppName(),
		Description: toPtr(gofakeit.Sentence(20)),
	})
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
