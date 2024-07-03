package fixture

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/volatiletech/null/v8"
)

// Factory An utility to facilitate the creation of mock data for testing purposes.
type Factory struct {
	t   *testing.T
	ctx context.Context
	db  *sql.DB
}

func NewFactory(t *testing.T, ctx context.Context, db *sql.DB) *Factory {
	return &Factory{
		t:   t,
		ctx: ctx,
		db:  db,
	}
}

func (f *Factory) NewOrganization() *Organization {
	return &Organization{
		factory: f,
		model: models.Organization{
			ID:        domain.NewID().String(),
			CreatedAt: time.Now().UTC(),
		},
	}
}

func (f *Factory) NewMember() *Member {
	return &Member{
		factory: f,
		model: models.Member{
			ID:              domain.NewID().String(),
			FirstName:       null.StringFrom(gofakeit.FirstName()),
			LastName:        null.StringFrom(gofakeit.LastName()),
			Email:           gofakeit.Email(),
			LoginProviderID: fmt.Sprintf("user_%s", domain.NewID().String()),
			OrganizationID:  "",
			CreatedAt:       time.Now().UTC(),
		},
	}
}

func (f *Factory) NewEnvironment() *Environment {
	return &Environment{
		factory: f,
		model: models.Environment{
			ID:             domain.NewID().String(),
			OrganizationID: domain.NewID().String(),
			Name:           gofakeit.AppName(),
			EnvType:        []string{models.EnvtypeDevelopment, models.EnvtypeProduction}[gofakeit.Number(0, 1)],
			CreatedAt:      time.Now().UTC(),
		},
	}
}

func (f *Factory) NewApplication() *Application {
	return &Application{
		factory: f,
		model: models.Application{
			ID:        domain.NewID().String(),
			Name:      gofakeit.AppName(),
			CreatedAt: time.Now().UTC(),
		},
	}
}

func (f *Factory) NewEventType() *EventType {
	return &EventType{
		factory: f,
		model: models.EventType{
			ID:   domain.NewID().String(),
			Name: gofakeit.Verb(),
		},
	}
}

func (f *Factory) NewApiKey() *ApiKey {
	ak, err := domain.NewApiKey(gofakeit.AppName(), domain.NewID(), domain.NewID(), nil, false)
	require.NoError(f.t, err)

	return &ApiKey{
		factory: f,
		model: models.APIKey{
			SecretKey: ak.SecretKey().FullKey(),
			Suffix:    ak.SecretKey().String(),
			Name:      ak.Name(),
		},
	}
}
