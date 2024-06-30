package psql_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/tests/fixture"
)

func TestApplicationRepository_Insert(t *testing.T) {
	ff := fixture.NewFactory(t, ctx, db)
	org := ff.NewOrganization().Save()
	env := ff.NewEnvironment().WithOrganizationID(org.ID).Save()
	app := ff.NewApplication().WithEnvironmentID(env.ID).NewDomainModel()

	require.NoError(t, applicationRepo.Insert(ctx, app))
}
