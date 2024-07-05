package command

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

const (
	maxResponseSize = 1024 * 1024 * 256
)

type CallWebhookEndpoint struct {
	EndpointID string
	MessageID  string
}

type CallWebhookEndpointHandler struct {
	txProvider TransactionProvider
}

func NewCallWebhookEndpointHandler(txProvider TransactionProvider) CallWebhookEndpointHandler {
	return CallWebhookEndpointHandler{
		txProvider: txProvider,
	}
}

func (c CallWebhookEndpointHandler) Execute(ctx context.Context, cmd CallWebhookEndpoint) error {
	endpointID, err := domain.NewIdFromString(cmd.EndpointID)
	if err != nil {
		return err
	}

	messageID, err := domain.NewIdFromString(cmd.MessageID)
	if err != nil {
		return err
	}

	return c.txProvider.Transact(ctx, func(adapters TransactableAdapters) error {
		endpoint, err := adapters.EndpointRepository.ByID(ctx, endpointID)
		if err != nil {
			return fmt.Errorf("error querying the endpoint by id '%s': %v", cmd.EndpointID, err)
		}

		message, err := adapters.MessageRepository.ByID(ctx, messageID)
		if err != nil {
			return fmt.Errorf("error querying message by id '%s': %v", cmd.MessageID, err)
		}

		httpClient := http.Client{
			Timeout: time.Second * 30,
		}

		// TODO: retrieve the notification from the endpoint itself
		req, err := http.NewRequestWithContext(
			ctx,
			http.MethodPost,
			endpoint.EndpointURL().String(),
			bytes.NewReader([]byte(message.Payload())),
		)
		if err != nil {
			return fmt.Errorf("error creating request: %v", err)
		}

		timestamp := time.Now()
		signature, err := createSignature(message, endpoint, timestamp)
		if err != nil {
			return err
		}

		req.Header.Set("user-agent", "subscribed-backend")
		req.Header.Set("x-sbs-id", cmd.MessageID)
		req.Header.Set("x-sbs-timestamp", fmt.Sprintf("%d", timestamp.Unix()))
		req.Header.Set("x-sbs-signature", signature)
		for name, value := range endpoint.Headers() {
			req.Header.Set(name, value)
		}

		resp, err := httpClient.Do(req)
		if err != nil {
			return fmt.Errorf("error making request to endpoint '%s': %v", endpoint.EndpointURL(), err)
		}
		defer func() { _ = resp.Body.Close() }()

		body := make([]byte, maxResponseSize)
		n, err := resp.Body.Read(body)
		if err != nil && !errors.Is(err, io.EOF) {
			return fmt.Errorf("error reading response from endpoint '%s': %v", endpoint.EndpointURL(), err)
		}

		reqHeaders := domain.Headers{}
		for name, value := range req.Header {
			reqHeaders[name] = strings.Join(value, ";")
		}

		attempt, err := domain.NewMessageSendAttempt(
			endpointID,
			messageID,
			string(body[:n]),
			domain.StatusCode(resp.StatusCode),
			reqHeaders,
		)
		if err != nil {
			return fmt.Errorf("error creating messageSendAttempt: %v", err)
		}

		err = adapters.MessageRepository.SaveMessageSendAttempt(ctx, attempt)
		if err != nil {
			return fmt.Errorf("error saving messageSendAttempt: %v", err)
		}

		return nil
	})
}

func createSignature(message *domain.Message, endpoint *domain.Endpoint, timestamp time.Time) (string, error) {
	signedContent := fmt.Sprintf("%s.%d.%s", message.Id(), timestamp.Unix(), message.Payload())

	// Decode secret (base64)
	secretBytes, err := endpoint.SigningSecret().UnEncoded()
	if err != nil {
		return "", fmt.Errorf("error decoding secret: %v", err)
	}

	// Create HMAC signature
	h := hmac.New(sha256.New, secretBytes)
	h.Write([]byte(signedContent))

	return base64.StdEncoding.EncodeToString(h.Sum(nil)), nil
}
