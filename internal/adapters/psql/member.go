package psql

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type MemberRepository struct {
	db boil.ContextExecutor
}

func NewMemberRepository(db boil.ContextExecutor) *MemberRepository {
	return &MemberRepository{
		db: db,
	}
}

func (o MemberRepository) Insert(ctx context.Context, member *iam.Member) error {
	model := models.Member{
		ID:              member.Id().String(),
		FirstName:       null.StringFrom(member.FirstName()),
		LastName:        null.StringFrom(member.LastName()),
		Email:           member.Email().String(),
		LoginProviderID: member.LoginProviderId().String(),
		OrganizationID:  member.OrganizationID().String(),
		CreatedAt:       member.CreatedAt(),
	}

	err := model.Insert(ctx, o.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}
