package fixture

import (
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
	"github.com/subscribeddotdev/subscribed-backend/tests"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type Member struct {
	factory           *Factory
	plainTextPassword string
	model             models.Member
}

func (m *Member) WithID(id string) *Member {
	m.model.ID = id
	return m
}

func (m *Member) WithFirstName(value string) *Member {
	m.model.FirstName = value
	return m
}

func (m *Member) WithLastName(value string) *Member {
	m.model.LastName = value
	return m
}

func (m *Member) WithEmail(value string) *Member {
	m.model.Email = value
	return m
}

func (m *Member) WithOrganizationID(value string) *Member {
	m.model.OrganizationID = value
	return m
}

func (m *Member) Save() (models.Member, string) {
	err := m.model.Insert(m.factory.ctx, m.factory.db, boil.Infer())
	require.NoError(m.factory.t, err)

	return m.model, m.plainTextPassword
}

func (m *Member) NewDomainModel() *iam.Member {
	member, err := iam.NewMember(
		iam.OrgID(m.model.OrganizationID),
		m.model.FirstName,
		m.model.LastName,
		tests.MustEmail(m.factory.t, m.model.Email),
		tests.FixturePassword(m.factory.t),
	)
	require.NoError(m.factory.t, err)
	return member
}
