package fixture

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Environment struct {
	factory *Factory
	model   models.Environment
}

func (e *Environment) WithID(id string) *Environment {
	e.model.ID = id
	return e
}

func (e *Environment) WithName(value string) *Environment {
	e.model.Name = value
	return e
}

func (e *Environment) WithEnvType(value string) *Environment {
	e.model.EnvType = value
	return e
}

func (e *Environment) WithArchivedAt(value time.Time) *Environment {
	e.model.ArchivedAt = null.TimeFrom(value)
	return e
}

func (e *Environment) WithOrganizationID(value string) *Environment {
	e.model.OrganizationID = value
	return e
}

func (e *Environment) Save() models.Environment {
	err := e.model.Insert(e.factory.ctx, e.factory.db, boil.Infer())
	require.NoError(e.factory.t, err)

	return e.model
}

func (e *Environment) NewDomainModel() *domain.Environment {
	environments := []struct {
		name    string
		envType domain.EnvType
	}{
		{
			name:    "Development",
			envType: domain.EnvTypeDevelopment,
		},
		{
			name:    "Production",
			envType: domain.EnvTypeProduction,
		},
	}

	env := environments[gofakeit.IntRange(0, len(environments)-1)]

	newEnv, err := domain.NewEnvironment(
		env.name,
		e.model.OrganizationID,
		env.envType,
	)
	require.NoError(e.factory.t, err)

	return newEnv
}
