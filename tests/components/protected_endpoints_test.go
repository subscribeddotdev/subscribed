package components_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	testjwks "github.com/subscribeddotdev/subscribed-backend/tests/jwks"
)

func TestProtectedEndpoints(t *testing.T) {
	token := testjwks.JwtGenerator(t, map[string]any{
		"sid": "sess_123",
		"sub": "user_123",
		"iss": "https://clerk.com",
	})
	resp, err := getClient(t, token).GetHelloWorld(ctx)
	require.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
