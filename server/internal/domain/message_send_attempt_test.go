package domain_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed/server/internal/domain"
)

func TestNewMessageSendAttempt(t *testing.T) {
	testCases := []struct {
		name           string
		expectedErr    string
		endpointID     domain.EndpointID
		messageID      domain.MessageID
		response       string
		statusCode     domain.StatusCode
		requestHeaders domain.Headers
	}{
		{
			name:        "create_new_message_send_attempt",
			expectedErr: "",
			endpointID:  domain.NewEndpointID(),
			messageID:   domain.NewMessageID(),
			response:    gofakeit.Sentence(10),
			statusCode:  domain.StatusCode(gofakeit.Number(200, 599)),
			requestHeaders: map[string]string{
				"UserAgent": gofakeit.AppName(),
			},
		},
		{
			name:        "error_empty_endpoint_id",
			expectedErr: "endpointID cannot be empty",
			endpointID:  domain.EndpointID(""),
			messageID:   domain.NewMessageID(),
			response:    gofakeit.Sentence(10),
			statusCode:  domain.StatusCode(gofakeit.Number(200, 599)),
			requestHeaders: map[string]string{
				"UserAgent": gofakeit.AppName(),
			},
		},
		{
			name:        "error_empty_message_id",
			expectedErr: "messageID cannot be empty",
			endpointID:  domain.NewEndpointID(),
			messageID:   domain.MessageID(""),
			response:    gofakeit.Sentence(10),
			statusCode:  domain.StatusCode(gofakeit.Number(200, 599)),
			requestHeaders: map[string]string{
				"UserAgent": gofakeit.AppName(),
			},
		},
		{
			name:        "error_invalid_status_code",
			expectedErr: "invalid status code",
			endpointID:  domain.NewEndpointID(),
			messageID:   domain.NewMessageID(),
			response:    gofakeit.Sentence(10),
			statusCode:  domain.StatusCode(gofakeit.Number(0, 199)),
			requestHeaders: map[string]string{
				"UserAgent": gofakeit.AppName(),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			attempt, err := domain.NewMessageSendAttempt(
				tc.endpointID,
				tc.messageID,
				tc.response,
				tc.statusCode,
				tc.requestHeaders,
			)
			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.endpointID, attempt.EndpointID())
			assert.Equal(t, tc.messageID, attempt.MessageID())
			assert.Equal(t, tc.response, attempt.Response())
			assert.Equal(t, tc.statusCode, attempt.StatusCode())
			assert.Len(t, attempt.RequestHeaders(), len(tc.requestHeaders))
			assert.Equal(t, tc.requestHeaders, attempt.RequestHeaders())

			if tc.statusCode < 300 {
				assert.Equal(t, "succeeded", attempt.Status().String())
			} else {
				assert.Equal(t, "failed", attempt.Status().String())
			}
		})
	}
}
