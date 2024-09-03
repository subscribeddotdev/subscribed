package fixture

import (
	"time"

	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type EventType struct {
	factory *Factory
	model   models.EventType
}

func (a *EventType) WithOrgID(value string) *EventType {
	a.model.OrgID = value
	return a
}

func (a *EventType) WithName(value string) *EventType {
	a.model.Name = value
	return a
}

func (a *EventType) WithDescription(value string) *EventType {
	a.model.Description = null.StringFrom(value)
	return a
}

func (a *EventType) WithSchema(value string) *EventType {
	a.model.Schema = null.StringFrom(value)
	return a
}

func (a *EventType) WithSchemaExample(value string) *EventType {
	a.model.SchemaExample = null.StringFrom(value)
	return a
}

func (a *EventType) WithArchivedAt(value time.Time) *EventType {
	a.model.ArchivedAt = null.TimeFrom(value)
	return a
}

func (a *EventType) Save() models.EventType {
	err := a.model.Insert(a.factory.ctx, a.factory.db, boil.Infer())
	require.NoError(a.factory.t, err)

	return a.model
}

func (a *EventType) NewDomainModel() *domain.EventType {
	return domain.UnMarshallEventType(
		domain.EventTypeID(a.model.ID),
		a.model.OrgID,
		a.model.Name,
		a.model.Description.String,
		a.model.Schema.String,
		a.model.SchemaExample.String,
		a.model.CreatedAt,
		a.model.ArchivedAt.Ptr(),
	)
}
