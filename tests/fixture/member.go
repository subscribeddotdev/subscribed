package fixture

import (
	"fmt"

	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
	"github.com/subscribeddotdev/subscribed-backend/tests"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Member struct {
	factory *Factory
	model   models.Member
}

func (m *Member) WithID(id string) *Member {
	m.model.ID = id
	return m
}

func (m *Member) WithFirstName(value string) *Member {
	m.model.FirstName = null.StringFrom(value)
	return m
}

func (m *Member) WithLastName(value string) *Member {
	m.model.LastName = null.StringFrom(value)
	return m
}

func (m *Member) WithEmail(value string) *Member {
	m.model.Email = value
	return m
}

func (m *Member) WithLoginProviderID(value string) *Member {
	m.model.LoginProviderID = fmt.Sprintf("user_%s", value)
	return m
}

func (m *Member) WithOrganizationID(value string) *Member {
	m.model.OrganizationID = value
	return m
}

func (m *Member) Save() models.Member {
	err := m.model.Insert(m.factory.ctx, m.factory.db, boil.Infer())
	require.NoError(m.factory.t, err)

	return m.model
}

func (m *Member) NewDomainModel() *iam.Member {
	member, err := iam.NewMember(
		iam.OrgID(m.model.OrganizationID),
		m.model.FirstName.String,
		m.model.LastName.String,
		tests.MustEmail(m.factory.t, m.model.Email),
		tests.FixturePassword(m.factory.t),
	)
	require.NoError(m.factory.t, err)
	return member
}
