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

		applicationID          domain.ID
		endpointUrl            domain.EndpointURL
		description            string
		eventTypesSubscribedTo []domain.EventType
	}{
		{
			name:                   "create_new_endpoint",
			expectedErr:            "",
			applicationID:          domain.NewID(),
			endpointUrl:            mustEndpointURL(t, gofakeit.URL()),
			description:            gofakeit.Sentence(gofakeit.IntRange(5, 10)),
			eventTypesSubscribedTo: nil,
		},
		{
			name:                   "error_invalid_endpoint_url",
			expectedErr:            "endpointURL cannot be empty",
			applicationID:          domain.NewID(),
			endpointUrl:            domain.EndpointURL{},
			description:            gofakeit.Sentence(gofakeit.IntRange(5, 10)),
			eventTypesSubscribedTo: nil,
		},
		{
			name:                   "error_invalid_application_id",
			expectedErr:            "applicationID cannot be empty",
			applicationID:          domain.ID{},
			endpointUrl:            mustEndpointURL(t, gofakeit.URL()),
			description:            gofakeit.Sentence(gofakeit.IntRange(5, 10)),
			eventTypesSubscribedTo: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			endpoint, err := domain.NewEndpoint(tc.endpointUrl, tc.applicationID, tc.description, tc.eventTypesSubscribedTo)

			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				return
			}

			require.NoError(t, err)

			assert.True(t, endpoint.CreatedAt().Before(time.Now()))
			assert.Equal(t, tc.endpointUrl, endpoint.EndpointURL())
			assert.Equal(t, tc.description, endpoint.Description())
			assert.Equal(t, tc.eventTypesSubscribedTo, endpoint.EventTypesSubscribedTo())
			assert.NotEmpty(t, endpoint.SigningSecret())
		})
	}
}

func mustEndpointURL(t *testing.T, rawURL string) domain.EndpointURL {
	e, err := domain.NewEndpointURL(rawURL)
	require.NoError(t, err)
	return e
}
