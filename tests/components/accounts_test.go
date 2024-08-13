package components_test

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/tests/client"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func TestSignup(t *testing.T) {
	apiClient := getClient(t, "")

	reqBody := client.SignupRequest{
		FirstName: gofakeit.FirstName(),
		LastName:  gofakeit.LastName(),
		Email:     gofakeit.Email(),
		Password:  gofakeit.Password(true, true, true, true, false, 12),
	}

	resp1, err := apiClient.SignUp(ctx, reqBody)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, resp1.StatusCode)
	assertMemberCreated(t, reqBody)
}

func assertMemberCreated(t *testing.T, reqBody client.SignupRequest) {
	member := findMemberByEmail(t, reqBody.Email)
	require.NotNil(t, member.R.Organization)
	assert.Equal(t, reqBody.FirstName, member.FirstName)
	assert.Equal(t, reqBody.LastName, member.LastName)
	assert.Equal(t, reqBody.Email, member.Email)
}

func findMemberByEmail(t *testing.T, email string) *models.Member {
	member, err := models.Members(
		models.MemberWhere.Email.EQ(email),
		qm.Load(models.MemberRels.Organization),
	).One(ctx, db)
	require.NoError(t, err)

	return member
}
