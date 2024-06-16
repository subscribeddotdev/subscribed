package psql_test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/tests"
)

func TestEnvironmentRepository_Insert(t *testing.T) {
	org := fixtureAndSaveOrganization(t)
	env0 := fixtureEnvironment(t, org.ID)
	err := environmentRepo.Insert(ctx, env0)
	require.NoError(t, err)
	assertEnvironment(t, env0)
}

func fixtureEnvironment(t *testing.T, orgID string) *domain.Environment {
	environments := []struct {
		name    string
		envType domain.EnvType
	}{
		{
			name:    "Development",
			envType: domain.EnvTypeDevelopment,
		}, {
			name:    "Production",
			envType: domain.EnvTypeProduction,
		},
	}

	env := environments[gofakeit.IntRange(0, len(environments)-1)]

	newEnv, err := domain.NewEnvironment(env.name, tests.MustID(t, orgID), env.envType)
	require.NoError(t, err)

	return newEnv
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
