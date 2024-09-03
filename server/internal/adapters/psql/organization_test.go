package psql_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

func TestOrganizationRepository_Insert(t *testing.T) {
	org := iam.NewOrganization()
	err := organizationRepo.Insert(ctx, org)
	require.NoError(t, err)
}
