package psql_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/psql"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
	"github.com/subscribeddotdev/subscribed-backend/tests"
	"github.com/subscribeddotdev/subscribed-backend/tests/fixture"
)

func TestMemberRepository_Lifecycle(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	repo := psql.NewMemberRepository(db)
	org := ff.NewOrganization().Save()
	member := ff.NewMember().WithOrganizationID(org.ID).NewDomainModel()

	t.Run("insert_new_member", func(t *testing.T) {
		require.NoError(t, repo.Insert(ctx, member))
	})

	t.Run("member_does_not_exist", func(t *testing.T) {
		_, err := repo.FindByEmail(ctx, tests.MustEmail(t, gofakeit.Email()))
		require.ErrorIs(t, err, iam.ErrMemberNotFound)
	})

	t.Run("find_member_by_email", func(t *testing.T) {
		foundMember, err := repo.FindByEmail(ctx, member.Email())
		require.NoError(t, err)
		assert.NotNil(t, foundMember)
	})
}
