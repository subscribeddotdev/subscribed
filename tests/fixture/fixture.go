package fixture

import (
	"context"
	"database/sql"
	"os"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
	"github.com/subscribeddotdev/subscribed-backend/tests"
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
			ID:        iam.NewOrgID().String(),
			CreatedAt: time.Now(),
		},
	}
}

func (f *Factory) NewMember() *Member {
	plainTextPassword := tests.FixturePassword(f.t).String()
	password, err := iam.NewPassword(plainTextPassword)
	require.NoError(f.t, err)

	return &Member{
		factory:           f,
		plainTextPassword: plainTextPassword,
		model: models.Member{
			ID:             iam.NewMemberID().String(),
			FirstName:      gofakeit.FirstName(),
			LastName:       gofakeit.LastName(),
			Email:          gofakeit.Email(),
			Password:       password.String(),
			OrganizationID: "",
			CreatedAt:      time.Now(),
		},
	}
}

func (f *Factory) NewEnvironment() *Environment {
	return &Environment{
		factory: f,
		model: models.Environment{
			ID:             domain.NewEnvironmentID().String(),
			OrganizationID: iam.NewOrgID().String(),
			Name:           gofakeit.AppName(),
			EnvType: []string{
				models.EnvtypeDevelopment,
				models.EnvtypeProduction,
			}[gofakeit.Number(0, 1)],
			CreatedAt: time.Now(),
		},
	}
}

func (f *Factory) NewApplication() *Application {
	return &Application{
		factory: f,
		model: models.Application{
			ID:        domain.NewApplicationID().String(),
			Name:      gofakeit.AppName(),
			CreatedAt: time.Now(),
		},
	}
}

func (f *Factory) NewEventType() *EventType {
	return &EventType{
		factory: f,
		model: models.EventType{
			ID:   domain.NewEventTypeID().String(),
			Name: gofakeit.Verb(),
		},
	}
}

func (f *Factory) NewApiKey() *ApiKey {
	ak, err := domain.NewApiKey(
		gofakeit.AppName(),
		iam.NewOrgID().String(),
		domain.NewEnvironmentID(),
		nil,
		false,
	)
	require.NoError(f.t, err)

	return &ApiKey{
		factory: f,
		model: models.APIKey{
			ID:        domain.NewApiKeyID().String(),
			SecretKey: ak.SecretKey().FullKey(),
			Suffix:    ak.SecretKey().String(),
			Name:      ak.Name(),
		},
	}
}

func (f *Factory) NewEndpoint() *Endpoint {
	ss, err := domain.NewSigningSecret()
	require.NoError(f.t, err)
	return &Endpoint{
		factory: f,
		model: models.Endpoint{
			ID:            domain.NewEndpointID().String(),
			ApplicationID: domain.NewID().String(),
			URL:           os.Getenv("WEBHOOK_EMULATOR_URL") + "/webhook",
			Description:   null.StringFrom(gofakeit.Sentence(10)),
			SigningSecret: ss.String(),
		},
	}
}
