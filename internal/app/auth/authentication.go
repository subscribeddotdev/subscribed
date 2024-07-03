package auth

import (
	"context"

	"github.com/subscribeddotdev/subscribed-backend/internal/domain"
	"github.com/subscribeddotdev/subscribed-backend/internal/domain/iam"
)

type Service struct {
	memberRepo iam.MemberRepository
	apiKeyRepo domain.ApiKeyRepository
}

func NewService(memberRepo iam.MemberRepository, apiKeyRepo domain.ApiKeyRepository) *Service {
	return &Service{
		memberRepo: memberRepo,
		apiKeyRepo: apiKeyRepo,
	}
}

func (s *Service) ResolveMemberByLoginProviderID(ctx context.Context, loginProviderID string) (*iam.Member, error) {
	return s.memberRepo.ByLoginProviderID(ctx, iam.LoginProviderID(loginProviderID))
}

func (s *Service) IsApiKeyValid(ctx context.Context, secretKey string) error {
	sk, err := domain.UnMarshallSecretKey(secretKey)
	if err != nil {
		return err
	}

	ak, err := s.apiKeyRepo.FindBySecretKey(ctx, sk)
	if err != nil {
		return err
	}

	if ak.IsExpired() {
		return domain.ErrApiKeyIsExpired
	}

	return nil
}
