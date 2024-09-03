package fixture

import (
	"time"

	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Organization struct {
	factory *Factory
	model   models.Organization
}

func (o *Organization) Save() models.Organization {
	err := o.model.Insert(o.factory.ctx, o.factory.db, boil.Infer())
	require.NoError(o.factory.t, err)

	return o.model
}

func (o *Organization) WithDisabledAt(disabledAt time.Time) *Organization {
	o.model.DisabledAt = null.TimeFrom(disabledAt)
	return o
}
