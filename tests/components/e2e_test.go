package components_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	"github.com/subscribeddotdev/subscribed-backend/tests/client"
)

func TestE2E(t *testing.T) {
	ctx := context.Background()
	publicClient := getClient(t)

	email := gofakeit.Email()
	password := uuid.NewString()

	signUpResp, err := publicClient.SignUp(ctx, client.SignUpJSONRequestBody{
		Email:     email,
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Password:  password,
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, signUpResp.StatusCode)

	signInResp, err := publicClient.SignInWithResponse(ctx, client.SignInJSONRequestBody{
		Email:    email,
		Password: password,
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, signInResp.StatusCode())

	token := signInResp.JSON200.Token
	authClient := getClientWithToken(t, token)

	createEventResp, err := authClient.CreateEventType(ctx, client.CreateEventTypeJSONRequestBody{
		Name: "order_placed",
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, createEventResp.StatusCode)

	createEventResp, err = authClient.CreateEventType(ctx, client.CreateEventTypeJSONRequestBody{
		Name: "order_refunded",
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, createEventResp.StatusCode)

	t.Run("should_not_create_an_app_with_token", func(t *testing.T) {
		resp, err := authClient.CreateApplication(ctx, client.CreateApplicationJSONRequestBody{
			Name: "Web App",
		})
		require.NoError(t, err)
		require.Equal(t, http.StatusForbidden, resp.StatusCode)
	})

	environmentsResp, err := authClient.GetEnvironmentsWithResponse(ctx)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, environmentsResp.StatusCode())

	var prodEnvID string
	for _, env := range environmentsResp.JSON200.Data {
		if env.Type == client.Production {
			prodEnvID = env.Id
			break
		}
	}
	require.NotEmpty(t, prodEnvID, "should find a production environment")

	apiKeyResp, err := authClient.CreateApiKeyWithResponse(ctx, client.CreateApiKeyJSONRequestBody{
		Name:          "test api key",
		EnvironmentId: prodEnvID,
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, apiKeyResp.StatusCode())

	apiKey := apiKeyResp.JSON201.UnmaskedApiKey

	apiKeyClient := getClientWithApiKey(t, apiKey)

	createAppResp, err := apiKeyClient.CreateApplicationWithResponse(ctx, client.CreateApplicationJSONRequestBody{
		Name: "Web App",
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, createAppResp.StatusCode())

	appID := createAppResp.JSON201.Id

	addEndpointResp, err := authClient.AddEndpoint(ctx, appID, client.AddEndpointJSONRequestBody{
		Description:  strPtr("All event types"),
		EventTypeIds: nil,
		Url:          "http://localhost:8090/webhook",
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, addEndpointResp.StatusCode)
}

func strPtr(s string) *string {
	return &s
}
