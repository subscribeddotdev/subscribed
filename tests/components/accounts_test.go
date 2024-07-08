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
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func TestAccountCreationWebhook(t *testing.T) {
	cli := getClient(t, "")

	params := &client.CreateAccountParams{
		SvixId:        fmt.Sprintf("msg_%s", gofakeit.UUID()),
		SvixTimestamp: fmt.Sprintf("%d", time.Now().Unix()),
		SvixSignature: gofakeit.UUID(),
	}
	reqBody := client.CreateAccountRequest{
		Data: client.ClerkWebhookUserCreatedData{
			Id:        fmt.Sprintf("user_%s", domain.NewID().String()),
			FirstName: toPtr(gofakeit.FirstName()),
			LastName:  toPtr(gofakeit.FirstName()),
			EmailAddresses: []client.ClerkWebhookEmailAddress{
				{
					EmailAddress: gofakeit.Email(),
					Id:           gofakeit.UUID(),
				},
			},
			CreatedAt: int(time.Now().Unix()),
		},
	}

	resp1, err := cli.CreateAccount(ctx, params, reqBody)
	require.NoError(t, err)

	assert.Equal(t, http.StatusCreated, resp1.StatusCode)
	assertMemberCreated(t, reqBody)
	assertOrgsDefaultEnvironments(t, reqBody.Data.Id)

	t.Run("create_account_should_be_idempotent", func(t *testing.T) {
		resp2, err := cli.CreateAccount(ctx, params, reqBody)
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp2.StatusCode)

		reqBody.Data.EmailAddresses[0].EmailAddress = gofakeit.Email()
		resp3, err := cli.CreateAccount(ctx, params, reqBody)
		require.NoError(t, err)
		assert.Equal(t, http.StatusCreated, resp3.StatusCode)

		assertMemberDuplication(t, reqBody.Data.Id)
		assertOrgsDefaultEnvironments(t, reqBody.Data.Id)
	})
}

func assertMemberCreated(t *testing.T, reqBody client.CreateAccountRequest) {
	member := findMemberByLoginProviderID(t, reqBody.Data.Id)
	require.NotNil(t, member.R.Organization)
	assert.Equal(t, reqBody.Data.FirstName, member.FirstName.Ptr())
	assert.Equal(t, reqBody.Data.LastName, member.LastName.Ptr())
	assert.Equal(t, reqBody.Data.EmailAddresses[0].EmailAddress, member.Email)
}

func assertOrgsDefaultEnvironments(t *testing.T, loginProviderID string) {
	member := findMemberByLoginProviderID(t, loginProviderID)
	envs, err := models.Environments(models.EnvironmentWhere.OrganizationID.EQ(member.OrganizationID)).All(ctx, db)
	require.NoError(t, err)

	require.Len(t, envs, 2, "two default environments should be created")
}

func assertMemberDuplication(t *testing.T, loginProviderId string) {
	totalMembersExpected := 1
	total, err := models.Members(models.MemberWhere.LoginProviderID.EQ(loginProviderId)).Count(ctx, db)
	require.NoError(t, err)
	assert.Equal(t, totalMembersExpected, int(total))
}

func findMemberByLoginProviderID(t *testing.T, loginProviderID string) *models.Member {
	member, err := models.Members(
		models.MemberWhere.LoginProviderID.EQ(loginProviderID),
		qm.Load(models.MemberRels.Organization),
	).One(ctx, db)
	require.NoError(t, err)

	return member
}
