package domain

import (
	"context"
)

type LoanRepository interface {
	Save(ctx context.Context, loan *Loan) error
	Update(ctx context.Context, id ID, fn func(loan *Loan) error) error
}

type MemberRepository interface {
	FindByID(ctx context.Context, id ID) (*Member, error)
}

type BookRepository interface {
	FindByID(ctx context.Context, id ID) (*Book, error)
	FindCopyByID(ctx context.Context, id ID) (*BookCopy, error)
	UpdateCopyAvailability(ctx context.Context, bc *BookCopy) error
}
