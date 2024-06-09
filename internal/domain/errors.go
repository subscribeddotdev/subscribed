package domain

import "errors"

var (
	ErrLoanNotFound          = errors.New("loan not found by id")
	ErrMemberNotFound        = errors.New("member not found by id")
	ErrBookNotFound          = errors.New("book not found by id")
	ErrBookCopyNotFound      = errors.New("book copy not found by id")
	ErrBookCopyAlreadyOnLoan = errors.New("the book copy is already on loan")
)
