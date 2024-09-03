package components_test

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed/server/tests/client"
	"github.com/subscribeddotdev/subscribed/server/tests/fixture"
)

func TestEventTypes(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	member, password := ff.NewMember().WithOrganizationID(org.ID).Save()
	token := signIn(t, member.Email, password)
	apiClient := getClientWithToken(t, token)

	resp, err := apiClient.CreateEventType(ctx, client.CreateEventTypeRequest{
		Name:        gofakeit.AppName(),
		Description: toPtr(gofakeit.Sentence(20)),
	})
	require.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
