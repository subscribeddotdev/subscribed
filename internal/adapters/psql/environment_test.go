package psql_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/tests/fixture"
)

func TestEnvironmentRepository_Insert(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).NewDomainModel()
	err := environmentRepo.Insert(ctx, env)

	require.NoError(t, err)
	assertEnvironment(t, env)
}

func assertEnvironment(t *testing.T, env *domain.Environment) {
	t.Helper()

	model, err := models.FindEnvironment(ctx, db, env.Id().String())
	require.NoError(t, err)

	assert.Equal(t, env.Name(), model.Name)
	assert.Equal(t, env.Type().String(), model.EnvType)
	assert.Nil(t, model.ArchivedAt.Ptr())
	assert.True(t, model.CreatedAt.UTC().Before(time.Now().UTC()))
}
