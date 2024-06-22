package psql_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
)

func TestApplicationRepository_Insert(t *testing.T) {
	org := fixtureAndSaveOrganization(t)
	env := fixtureEnvironment(t, org.ID)
	app := fixtureApplication(t, env.Id())

	require.NoError(t, environmentRepo.Insert(ctx, env))
	require.NoError(t, applicationRepo.Insert(ctx, app))
}

func fixtureApplication(t *testing.T, envID domain.ID) *domain.Application {
	a, err := domain.NewApplication(gofakeit.Company(), envID)
	require.NoError(t, err)
	return a
}
