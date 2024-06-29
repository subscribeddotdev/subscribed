package auth

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

type Service struct {
	memberRepo iam.MemberRepository
}

func NewService(memberRepo iam.MemberRepository) *Service {
	return &Service{memberRepo: memberRepo}
}

func (s *Service) ResolveMemberByLoginProviderID(ctx context.Context, loginProviderID string) (*iam.Member, error) {
	return s.memberRepo.ByLoginProviderID(ctx, iam.LoginProviderID(loginProviderID))
}
