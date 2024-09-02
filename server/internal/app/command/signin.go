package command

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

type SignIn struct {
	Email             iam.Email
	PlainTextPassword string
}

type SignInHandler struct {
	repo iam.MemberRepository
}

func NewSignInHandler(repo iam.MemberRepository) SignInHandler {
	return SignInHandler{
		repo: repo,
	}
}

func (c SignInHandler) Execute(ctx context.Context, cmd SignIn) (*iam.Member, error) {
	member, err := c.repo.FindByEmail(ctx, cmd.Email)
	if err != nil {
		return nil, err
	}

	err = member.Authenticate(cmd.PlainTextPassword)
	if err != nil {
		return nil, err
	}

	return member, nil
}
