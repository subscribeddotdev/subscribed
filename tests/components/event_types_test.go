package components_test

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/tests/client"
	"github.com/subscribeddotdev/subscribed-backend/tests/fixture"
)

func TestEventTypes(t *testing.T) {
	factory := fixture.NewFactory(t, ctx, db)
	org := factory.NewOrganization().Save()
	factory.NewMember().WithOrganizationID(org.ID).Save()
	token := "" // jwks.JwtGenerator(t, member.LoginProviderID)

	resp, err := getClient(t, token).CreateEventType(ctx, client.CreateEventTypeRequest{
		Name:        gofakeit.AppName(),
		Description: toPtr(gofakeit.Sentence(20)),
	})
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
