package components_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/tests/client"
	"github.com/subscribeddotdev/subscribed-backend/tests/jwks"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestEventTypes(t *testing.T) {
	orgModel := models.Organization{
		ID: domain.NewID().String(),
	}
	require.NoError(t, orgModel.Insert(ctx, db, boil.Infer()))

	memberModel := models.Member{
		ID:              domain.NewID().String(),
		FirstName:       null.StringFrom(gofakeit.FirstName()),
		LastName:        null.StringFrom(gofakeit.LastName()),
		Email:           gofakeit.Email(),
		LoginProviderID: fmt.Sprintf("user_%s", domain.NewID().String()),
		OrganizationID:  orgModel.ID,
		CreatedAt:       time.Now(),
	}
	require.NoError(t, memberModel.Insert(ctx, db, boil.Infer()))

	token := jwks.JwtGenerator(t, map[string]any{
		"sid": "sess_123",
		"sub": memberModel.LoginProviderID,
		"iss": "https://clerk.com",
	})

	resp, err := getClient(t, token).CreateEventType(ctx, client.CreateEventTypeRequest{
		Name:        gofakeit.AppName(),
		Description: toPtr(gofakeit.Sentence(20)),
	})
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
