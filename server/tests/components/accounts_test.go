package components_test

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/tests/client"
)

func TestSignup(t *testing.T) {
	apiClient := getClientWithToken(t, noToken)

	req := client.SignupRequest{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
		Password:  gofakeit.Password(true, true, true, true, false, 12),
	}
	resp1, err := apiClient.SignUp(ctx, req)
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp1.StatusCode)

	token := signIn(t, req.Email, req.Password)
	require.NotEmpty(t, token)
}

func signIn(t *testing.T, email, password string) string {
	resp, err := getClientWithToken(t, noToken).SignInWithResponse(ctx, client.SigninRequest{
		Email:    email,
		Password: password,
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode())

	return resp.JSON200.Token
}
