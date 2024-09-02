package domain_test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

func TestNewEndpoint(t *testing.T) {
	testCases := []struct {
		name        string
		expectedErr string

		applicationID domain.ApplicationID
		endpointUrl   domain.EndpointURL
		description   string
		eventTypeIDs  []domain.EventTypeID
	}{
		{
			name:          "create_new_endpoint",
			expectedErr:   "",
			applicationID: domain.NewApplicationID(),
			endpointUrl:   mustEndpointURL(t, gofakeit.URL()),
			description:   gofakeit.Sentence(gofakeit.IntRange(5, 10)),
			eventTypeIDs:  []domain.EventTypeID{domain.NewEventTypeID()},
		},
		{
			name:          "error_invalid_endpoint_url",
			expectedErr:   "endpointURL cannot be empty",
			applicationID: domain.NewApplicationID(),
			endpointUrl:   domain.EndpointURL{},
			description:   gofakeit.Sentence(gofakeit.IntRange(5, 10)),
			eventTypeIDs:  []domain.EventTypeID{domain.NewEventTypeID()},
		},
		{
			name:          "error_invalid_application_id",
			expectedErr:   "applicationID cannot be empty",
			applicationID: domain.ApplicationID(""),
			endpointUrl:   mustEndpointURL(t, gofakeit.URL()),
			description:   gofakeit.Sentence(gofakeit.IntRange(5, 10)),
			eventTypeIDs:  []domain.EventTypeID{domain.NewEventTypeID()},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			endpoint, err := domain.NewEndpoint(tc.endpointUrl, tc.applicationID, tc.description, tc.eventTypeIDs)

			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				return
			}

			require.NoError(t, err)

			assert.True(t, endpoint.CreatedAt().Before(time.Now()))
			assert.Equal(t, tc.endpointUrl, endpoint.EndpointURL())
			assert.Equal(t, tc.description, endpoint.Description())
			assert.Equal(t, tc.eventTypeIDs, endpoint.EventTypeIDs())
			assert.NotEmpty(t, endpoint.SigningSecret())
		})
	}
}

func mustEndpointURL(t *testing.T, rawURL string) domain.EndpointURL {
	e, err := domain.NewEndpointURL(rawURL)
	require.NoError(t, err)
	return e
}
