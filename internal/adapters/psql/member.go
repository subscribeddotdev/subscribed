package psql

import (
	"context"
	"fmt"

	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

func (o MemberRepository) ExistsByOr(ctx context.Context, email iam.Email, loginProviderID iam.LoginProviderID) (bool, error) {
	exists, err := models.Members(
		models.MemberWhere.Email.EQ(email.String()),
		qm.Or2(models.MemberWhere.LoginProviderID.EQ(loginProviderID.String())),
	).Exists(ctx, o.db)
	if err != nil {
		return false, fmt.Errorf("unable to check for member via email '%s' or login provider id '%s': %v", email, loginProviderID, err)
	}

	return exists, nil
}
