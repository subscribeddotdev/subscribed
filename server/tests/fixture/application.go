package fixture

import (
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed/server/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed/server/internal/domain"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Application struct {
	factory *Factory
	model   models.Application
}

func (a *Application) WithEnvironmentID(value string) *Application {
	a.model.EnvironmentID = value
	return a
}

func (a *Application) WithName(value string) *Application {
	a.model.Name = value
	return a
}

func (a *Application) Save() models.Application {
	err := a.model.Insert(a.factory.ctx, a.factory.db, boil.Infer())
	require.NoError(a.factory.t, err)

	return a.model
}

func (a *Application) NewDomainModel() *domain.Application {
	app, err := domain.NewApplication(
		a.model.Name,
		domain.EnvironmentID(a.model.EnvironmentID),
	)
	require.NoError(a.factory.t, err)
	return app
}
