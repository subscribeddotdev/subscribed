package domain

import "time"

type Loan struct {
	id         ID
	loanedTo   *Member
	bookCopy   *BookCopy
	returnDate *time.Time
	startDate  time.Time
	dueDate    time.Time
}

func NewLoan(
	id ID,
	loanedTo *Member,
	bookCopy *BookCopy,
	returnDate *time.Time,
	startDate,
	dueDate time.Time,
) (*Loan, error) {
	return &Loan{
		id:         id,
		loanedTo:   loanedTo,
		bookCopy:   bookCopy,
		returnDate: returnDate,
		startDate:  startDate,
		dueDate:    dueDate,
	}, nil
}

func (l *Loan) ID() ID {
	return l.id
}

func (l *Loan) LoanedTo() *Member {
	return l.loanedTo
}

func (l *Loan) BookCopy() *BookCopy {
	return l.bookCopy
}

func (l *Loan) ReturnDate() *time.Time {
	return l.returnDate
}

func (l *Loan) StartDate() time.Time {
	return l.startDate
}

func (l *Loan) DueDate() time.Time {
	return l.dueDate
}

func (l *Loan) BookCopyReturned() {
	returnDate := time.Now()
	l.returnDate = &returnDate
	l.bookCopy.availability = BookAvailabilityAvailable
}

func (l *Loan) BookCopyLost() {
	l.bookCopy.availability = BookAvailabilityLost
}
