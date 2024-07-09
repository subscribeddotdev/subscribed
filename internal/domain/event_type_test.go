package domain_test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

func TestNewEventType(t *testing.T) {
	testCases := []struct {
		name        string
		expectedErr string

		orgID         string
		eventTypeName string
		description   string
		schema        string
		schemaExample string
	}{
		{
			name:          "create_new_event_type",
			expectedErr:   "",
			orgID:         iam.NewOrgID().String(),
			eventTypeName: gofakeit.VerbAction(),
			description:   gofakeit.Sentence(10),
			schema:        "",
			schemaExample: "",
		},
		{
			name:          "error_empty_or_missing_org_id",
			expectedErr:   "orgID cannot be empty",
			orgID:         " ",
			eventTypeName: gofakeit.VerbAction(),
			description:   gofakeit.Sentence(10),
			schema:        "",
			schemaExample: "",
		},
		{
			name:          "error_empty_event_type_name",
			expectedErr:   "name cannot be empty",
			orgID:         iam.NewOrgID().String(),
			eventTypeName: " ",
			description:   gofakeit.Sentence(10),
			schema:        "",
			schemaExample: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			env, err := domain.NewEventType(
				tc.orgID,
				tc.eventTypeName,
				tc.description,
				tc.schema,
				tc.schemaExample,
			)

			if tc.expectedErr != "" {
				assert.EqualError(t, err, tc.expectedErr)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tc.orgID, env.OrgID())
			assert.Equal(t, tc.eventTypeName, env.Name())
			assert.Equal(t, tc.description, env.Description())
			assert.Equal(t, tc.schema, env.Schema())
			assert.Equal(t, tc.schemaExample, env.SchemaExample())
			assert.True(t, env.CreatedAt().Before(time.Now()))
		})
	}
}
