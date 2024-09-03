package tests

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

func MustEmail(t *testing.T, address string) iam.Email {
	email, err := iam.NewEmail(address)
	require.NoError(t, err)
	return email
}

func FixturePassword(t *testing.T) iam.Password {
	p, err := iam.NewPassword(gofakeit.Password(true, true, true, true, false, 12))
	require.NoError(t, err)
	return p
}

func MustID(t *testing.T, id string) domain.ID {
	ID, err := domain.NewIdFromString(id)
	require.NoError(t, err)
	return ID
}

func ToPtr[T any](v T) *T {
	return &v
}

func RequireEqualTime(t *testing.T, t1, t2 time.Time) {
	t.Helper()
	require.Equal(t, t1.UTC().Truncate(time.Second), t2.Truncate(time.Second))
}
