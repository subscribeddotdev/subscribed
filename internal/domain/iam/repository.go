package iam

import (
	"context"
	"errors"
)

var (
	ErrMemberNotFound = errors.New("member not found")
)

type OrganizationRepository interface {
	Insert(ctx context.Context, org *Organization) error
}

type MemberRepository interface {
	Insert(ctx context.Context, member *Member) error
	FindByEmail(ctx context.Context, email Email) (*Member, error)
}
