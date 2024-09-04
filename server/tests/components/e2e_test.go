package components_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
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

	ts := NewTestServer()
	defer ts.Close()

	addEndpointResp, err := authClient.AddEndpoint(ctx, appID, client.AddEndpointJSONRequestBody{
		Description:  strPtr("All event types"),
		EventTypeIds: nil,
		Url:          ts.server.URL + "/all-event-types",
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, addEndpointResp.StatusCode)

	addEndpointResp, err = authClient.AddEndpoint(ctx, appID, client.AddEndpointJSONRequestBody{
		Description:  strPtr("Order placed only"),
		EventTypeIds: &[]string{orderPlacedEventTypeID},
		Url:          ts.server.URL + "/order-placed-only",
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, addEndpointResp.StatusCode)

	addEndpointResp, err = authClient.AddEndpoint(ctx, appID, client.AddEndpointJSONRequestBody{
		Description:  strPtr("Order refunded only"),
		EventTypeIds: &[]string{orderRefundedEventTypeID},
		Url:          ts.server.URL + "/order-refunded-only",
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusNoContent, addEndpointResp.StatusCode)

	addEndpointResp, err = authClient.AddEndpoint(ctx, appID, client.AddEndpointJSONRequestBody{
		Description:  strPtr("Both order placed and refunded"),
		EventTypeIds: &[]string{orderPlacedEventTypeID, orderRefundedEventTypeID},
		Url:          ts.server.URL + "/both-order-placed-and-refunded",
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

	resp, err := authClient.SendMessage(ctx, appID, client.SendMessageJSONRequestBody{
		EventTypeId: orderPlacedEventTypeID,
		Payload:     string(orderPlacedJSON),
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	resp, err = authClient.SendMessage(ctx, appID, client.SendMessageJSONRequestBody{
		EventTypeId: orderRefundedEventTypeID,
		Payload:     string(orderRefundedJSON),
	})
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, resp.StatusCode)

	require.EventuallyWithT(t, func(t *assert.CollectT) {
		assert.Equal(t, []string{orderPlaced.OrderID}, ts.orderPlacedOnly)
	}, 5*time.Second, 1*time.Second)

	require.EventuallyWithT(t, func(t *assert.CollectT) {
		assert.Equal(t, []string{orderRefunded.OrderID}, ts.orderRefundedOnly)
	}, 5*time.Second, 1*time.Second)

	require.EventuallyWithT(t, func(t *assert.CollectT) {
		assert.Equal(t, []string{orderPlaced.OrderID, orderRefundedEventTypeID}, ts.allEventTypes)
	}, 5*time.Second, 1*time.Second)

	require.EventuallyWithT(t, func(t *assert.CollectT) {
		assert.Equal(t, []string{orderPlaced.OrderID, orderRefundedEventTypeID}, ts.bothOrderPlacedAndRefunded)
	}, 5*time.Second, 1*time.Second)
}

type TestWebhookServer struct {
	server *httptest.Server

	allEventTypes              []string
	orderPlacedOnly            []string
	orderRefundedOnly          []string
	bothOrderPlacedAndRefunded []string
}

type Header struct {
	EventName string `json:"event_name"`
}

type Event struct {
	Header Header `json:"header"`
}

type OrderPlaced struct {
	Header  Header `json:"header"`
	OrderID string `json:"order_id"`
}

type OrderRefunded struct {
	Header  Header `json:"header"`
	OrderID string `json:"order_id"`
}

func NewTestServer() *TestWebhookServer {
	ts := &TestWebhookServer{}

	e := echo.New()
	e.POST("/all-event-types", func(c echo.Context) error {
		var event Event
		err := c.Bind(&event)
		if err != nil {
			return err
		}

		switch event.Header.EventName {
		case "order_placed":
			var orderPlaced OrderPlaced
			err := c.Bind(&orderPlaced)
			if err != nil {
				return err
			}

			ts.orderPlacedOnly = append(ts.orderPlacedOnly, orderPlaced.OrderID)
		case "order_refunded":
			var orderRefunded OrderRefunded
			err := c.Bind(&orderRefunded)
			if err != nil {
				return err
			}

			ts.orderRefundedOnly = append(ts.orderRefundedOnly, orderRefunded.OrderID)
		default:
			return c.NoContent(http.StatusBadRequest)
		}

		return c.NoContent(http.StatusOK)
	})
	e.POST("/order-placed-only", func(c echo.Context) error {
		var orderPlaced OrderPlaced
		err := c.Bind(&orderPlaced)
		if err != nil {
			return err
		}

		ts.orderPlacedOnly = append(ts.orderPlacedOnly, orderPlaced.OrderID)

		return c.NoContent(http.StatusOK)
	})
	e.POST("/order-refunded-only", func(c echo.Context) error {
		var orderRefunded OrderRefunded
		err := c.Bind(&orderRefunded)
		if err != nil {
			return err
		}

		ts.orderRefundedOnly = append(ts.orderRefundedOnly, orderRefunded.OrderID)

		return c.NoContent(http.StatusOK)
	})
	e.POST("/both-order-placed-and-refunded", func(c echo.Context) error {
		var event Event
		err := c.Bind(&event)
		if err != nil {
			return err
		}

		switch event.Header.EventName {
		case "order_placed":
			var orderPlaced OrderPlaced
			err := c.Bind(&orderPlaced)
			if err != nil {
				return err
			}

			ts.orderPlacedOnly = append(ts.orderPlacedOnly, orderPlaced.OrderID)
		case "order_refunded":
			var orderRefunded OrderRefunded
			err := c.Bind(&orderRefunded)
			if err != nil {
				return err
			}

			ts.orderRefundedOnly = append(ts.orderRefundedOnly, orderRefunded.OrderID)
		default:
			return c.NoContent(http.StatusBadRequest)
		}

		return c.NoContent(http.StatusOK)
	})

	ts.server = httptest.NewServer(e)

	return ts
}

func (t *TestWebhookServer) Close() {
	t.server.Close()
}

func strPtr(s string) *string {
	return &s
}
