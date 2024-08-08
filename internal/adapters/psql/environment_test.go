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

	model, err := models.FindEnvironment(ctx, db, env.ID().String())
	require.NoError(t, err)

	assertEnvironment(t, model, env)
}

func TestEnvironmentRepository_ByID(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()

	foundEnv, err := environmentRepo.ByID(ctx, domain.EnvironmentID(env.ID))
	require.NoError(t, err)

	assertEnvironment(t, &env, foundEnv)
}

func TestNewEnvironmentRepository_FindAll(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()

	envs := make(map[string]models.Environment)
	for i := 0; i < 3; i++ {
		e := ff.NewEnvironment().WithOrganizationID(org.ID).Save()
		envs[e.ID] = e
	}

	foundEnvs, err := environmentRepo.FindAll(ctx, org.ID)
	require.NoError(t, err)

	for _, foundEnv := range foundEnvs {
		env, exists := envs[foundEnv.ID().String()]
		require.True(t, exists)
		assertEnvironment(t, &env, foundEnv)
	}
}

func assertEnvironment(t *testing.T, model *models.Environment, env *domain.Environment) {
	t.Helper()

	assert.Equal(t, env.ID().String(), model.ID)
	assert.Equal(t, env.OrgID(), model.OrganizationID)
	assert.Equal(t, env.Name(), model.Name)
	assert.Equal(t, env.Type().String(), model.EnvType)
	assert.True(t, model.CreatedAt.Before(time.Now()))

	if model.ArchivedAt.Ptr() != nil {
		assert.Nil(t, env.ArchivedAt(), model.ArchivedAt.Ptr())
	}
}
