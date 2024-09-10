package components_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/subscribeddotdev/subscribed/server/tests/client"
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

	createEventResp, err := authClient.CreateEventTypeWithResponse(ctx, client.CreateEventTypeJSONRequestBody{
		Name: "order_placed",
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, createEventResp.StatusCode())

	orderPlacedEventTypeID := createEventResp.JSON201.Id

	createEventResp, err = authClient.CreateEventTypeWithResponse(ctx, client.CreateEventTypeJSONRequestBody{
		Name: "order_refunded",
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, createEventResp.StatusCode())

	orderRefundedEventTypeID := createEventResp.JSON201.Id

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
		Url:          webhookTargetURL(appID, "/all-event-types"),
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, addEndpointResp.StatusCode)

	addEndpointResp, err = authClient.AddEndpoint(ctx, appID, client.AddEndpointJSONRequestBody{
		Description:  strPtr("Order placed only"),
		EventTypeIds: &[]string{orderPlacedEventTypeID},
		Url:          webhookTargetURL(appID, "/order-placed-only"),
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, addEndpointResp.StatusCode)

	addEndpointResp, err = authClient.AddEndpoint(ctx, appID, client.AddEndpointJSONRequestBody{
		Description:  strPtr("Order refunded only"),
		EventTypeIds: &[]string{orderRefundedEventTypeID},
		Url:          webhookTargetURL(appID, "/order-refunded-only"),
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, addEndpointResp.StatusCode)

	addEndpointResp, err = authClient.AddEndpoint(ctx, appID, client.AddEndpointJSONRequestBody{
		Description:  strPtr("Both order placed and refunded"),
		EventTypeIds: &[]string{orderPlacedEventTypeID, orderRefundedEventTypeID},
		Url:          webhookTargetURL(appID, "/both-order-placed-and-refunded"),
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, addEndpointResp.StatusCode)

	orderPlaced := OrderPlaced{
		Header: Header{
			EventName: "order_placed",
		},
		OrderID: uuid.NewString(),
	}

	orderRefunded := OrderRefunded{
		Header: Header{
			EventName: "order_refunded",
		},
		OrderID: uuid.NewString(),
	}

	orderPlacedJSON, err := json.Marshal(orderPlaced)
	require.NoError(t, err)

	orderRefundedJSON, err := json.Marshal(orderRefunded)
	require.NoError(t, err)

	resp, err := apiKeyClient.SendMessage(ctx, appID, client.SendMessageJSONRequestBody{
		EventTypeId: orderPlacedEventTypeID,
		Payload:     string(orderPlacedJSON),
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp.StatusCode)

	resp, err = apiKeyClient.SendMessage(ctx, appID, client.SendMessageJSONRequestBody{
		EventTypeId: orderRefundedEventTypeID,
		Payload:     string(orderRefundedJSON),
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusCreated, resp.StatusCode)

	require.EventuallyWithT(t, func(t *assert.CollectT) {
		ids := getOrderIDFromWebhooks(t, appID, "/all-event-types")
		assert.Equal(t, []string{orderPlaced.OrderID, orderRefunded.OrderID}, ids)
	}, 2*time.Second, 200*time.Millisecond)

}

type Event struct {
	OrderID string `json:"order_id"`
}

type Header struct {
	EventName string `json:"event_name"`
}

type OrderPlaced struct {
	Header  Header `json:"header"`
	OrderID string `json:"order_id"`
}

type OrderRefunded struct {
	Header  Header `json:"header"`
	OrderID string `json:"order_id"`
}

func strPtr(s string) *string {
	return &s
}

func webhookTargetURL(id string, path string) string {
	return fmt.Sprintf("http://webhook-target:8080/webhooks/%s%s", id, path)
}

func localWebhookTargetURL(id string, path string) string {
	return fmt.Sprintf("http://localhost:8081/webhooks/%s%s", id, path)
}

// Note: Don't use require here because it's being used in EventuallyWithT
func getOrderIDFromWebhooks(t assert.TestingT, id string, path string) []string {
	resp, err := http.Get(localWebhookTargetURL(id, path))
	if !assert.NoError(t, err) {
		return nil
	}
	if !assert.Equal(t, http.StatusOK, resp.StatusCode) {
		return nil
	}

	var webhooks []string
	err = json.NewDecoder(resp.Body).Decode(&webhooks)
	if !assert.NoError(t, err) {
		return nil
	}

	var ids []string
	for _, webhook := range webhooks {
		var header Event
		err := json.Unmarshal([]byte(webhook), &header)
		if !assert.NoError(t, err) {
			return nil
		}
		ids = append(ids, header.OrderID)
	}

	return ids
}
