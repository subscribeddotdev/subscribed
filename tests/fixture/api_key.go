package fixture

import (
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type ApiKey struct {
	factory *Factory
	model   models.APIKey
}

func (a *ApiKey) WithEnvironmentID(value string) *ApiKey {
	a.model.EnvironmentID = value
	return a
}

func (a *ApiKey) Save() models.APIKey {
	err := a.model.Insert(a.factory.ctx, a.factory.db, boil.Infer())
	require.NoError(a.factory.t, err)

	return a.model
}
