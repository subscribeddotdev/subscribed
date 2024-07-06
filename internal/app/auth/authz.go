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

func (s *Service) resolveApiKeyFromSecretKey(ctx context.Context, secretKey string) (*domain.ApiKey, error) {
	sk, err := domain.UnMarshallSecretKey(secretKey)
	if err != nil {
		return nil, err
	}

	ak, err := s.apiKeyRepo.FindBySecretKey(ctx, sk)
	if err != nil {
		return nil, err
	}

	return ak, nil
}

func (s *Service) ResolveOrgIdFromSecretKey(ctx context.Context, secretKey string) (domain.ID, error) {
	ak, err := s.resolveApiKeyFromSecretKey(ctx, secretKey)
	if err != nil {
		return "", err
	}

	if ak.IsExpired() {
		return "", domain.ErrApiKeyIsExpired
	}

	return ak.OrgID(), nil
}
