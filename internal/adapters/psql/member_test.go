package psql_test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/psql"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
	"github.com/subscribeddotdev/subscribed-backend/tests"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestMemberRepository_Insert(t *testing.T) {
	repo := psql.NewMemberRepository(db)
	org := fixtureAndSaveOrganization(t)

	t.Run("insert_new_member", func(t *testing.T) {
		member := fixtureMember(t, org.ID)
		err := repo.Insert(ctx, member)

		require.NoError(t, err)
		assertMemberExists(t, member.LoginProviderId().String())
	})
}

func TestMemberRepository_ExistsByOr(t *testing.T) {
	repo := psql.NewMemberRepository(db)

	t.Run("member_does_not_exist", func(t *testing.T) {
		exists, err := repo.ExistsByOr(ctx, tests.MustEmail(t, gofakeit.Email()), iam.LoginProviderID(gofakeit.UUID()))
		require.NoError(t, err)
		assert.False(t, exists)
	})

	org := fixtureAndSaveOrganization(t)
	member := fixtureMember(t, org.ID)
	require.NoError(t, repo.Insert(ctx, member))

	t.Run("find_member_by_email", func(t *testing.T) {
		// When
		exists, err := repo.ExistsByOr(ctx, member.Email(), iam.LoginProviderID(gofakeit.UUID()))

		// Then
		require.NoError(t, err)
		assert.True(t, exists)
	})

	t.Run("find_member_by_login_provider_id", func(t *testing.T) {
		// When
		exists, err := repo.ExistsByOr(ctx, tests.MustEmail(t, gofakeit.Email()), member.LoginProviderId())

		// Then
		require.NoError(t, err)
		assert.True(t, exists)
	})
}

func fixtureAndSaveOrganization(t *testing.T) *models.Organization {
	model := &models.Organization{
		CreatedAt: time.Now(),
		ID:        domain.NewID().String(),
	}
	require.NoError(t, model.Insert(ctx, db, boil.Infer()))
	return model
}

func fixtureMember(t *testing.T, orgId string) *iam.Member {
	member, err := iam.NewMember(
		tests.MustID(t, orgId),
		iam.LoginProviderID(gofakeit.UUID()),
		gofakeit.FirstName(),
		gofakeit.LastName(),
		tests.MustEmail(t, gofakeit.Email()),
	)
	require.NoError(t, err)
	return member
}

func assertMemberExists(t *testing.T, loginProviderId string) {
	exists, err := models.Members(
		models.MemberWhere.LoginProviderID.EQ(loginProviderId),
	).Exists(ctx, db)
	require.NoError(t, err)
	assert.True(t, exists)
}
