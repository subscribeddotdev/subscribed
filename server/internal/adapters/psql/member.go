package psql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/friendsofgo/errors"
	"github.com/subscribeddotdev/subscribed-backend/internal/adapters/models"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
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
	foundMember, err := o.FindByEmail(ctx, member.Email())
	if err != nil && !errors.Is(err, iam.ErrMemberNotFound) {
		return fmt.Errorf("error checking for '%s' before creating an account: %v", member.Email(), err)
	}

	if foundMember != nil {
		return iam.ErrMemberAlreadyExists
	}

	model := models.Member{
		ID:             member.ID().String(),
		FirstName:      member.FirstName(),
		LastName:       member.LastName(),
		Email:          member.Email().String(),
		Password:       member.Password().String(),
		OrganizationID: member.OrgID().String(),
		CreatedAt:      member.CreatedAt(),
	}

	err = model.Insert(ctx, o.db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (o MemberRepository) FindByEmail(
	ctx context.Context,
	email iam.Email,
) (*iam.Member, error) {
	row, err := models.Members(models.MemberWhere.Email.EQ(email.String())).One(ctx, o.db)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, iam.ErrMemberNotFound
	}
	if err != nil {
		return nil, fmt.Errorf("unable to check for member via email '%s': %v", email, err)
	}

	return iam.UnMarshallMember(
		iam.MemberID(row.ID),
		iam.OrgID(row.OrganizationID),
		row.FirstName,
		row.LastName,
		row.Email,
		row.Password,
		row.CreatedAt,
	)
}
