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

func (s *Service) ResolveApiKeyFromSecretKey(ctx context.Context, secretKey string) (*domain.ApiKey, error) {
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
