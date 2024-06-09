package domain

import (
	"fmt"
	"time"
)

type Book struct {
	id              ID
	title           string
	subtitle        string
	publisher       string
	publicationDate time.Time
	authors         string
	isbn            ISBN
	format          BookFormat
	copies          []BookCopy
}

func NewBook(
	id ID,
	title,
	subtitle,
	publisher string,
	publicationDate time.Time,
	authors string,
	isbn ISBN,
	format BookFormat,
) (*Book, error) {
	return &Book{
		id:              id,
		title:           title,
		subtitle:        subtitle,
		publisher:       publisher,
		publicationDate: publicationDate,
		authors:         authors,
		isbn:            isbn,
		format:          format,
	}, nil
}

func (b *Book) ID() ID {
	return b.id
}

func (b *Book) Title() string {
	return b.title
}

func (b *Book) Subtitle() string {
	return b.subtitle
}

func (b *Book) Publisher() string {
	return b.publisher
}

func (b *Book) PublicationDate() time.Time {
	return b.publicationDate
}

func (b *Book) Authors() string {
	return b.authors
}

func (b *Book) Isbn() ISBN {
	return b.isbn
}

func (b *Book) Format() BookFormat {
	return b.format
}

func (b *Book) IsAvailable() bool {
	for _, bookCopy := range b.copies {
		if bookCopy.availability == BookAvailabilityAvailable {
			return true
		}
	}

	return false
}

type BookCopy struct {
	id           ID
	availability BookAvailability
}

func (b *BookCopy) ID() ID {
	return b.id
}

func (b *BookCopy) Availability() BookAvailability {
	return b.availability
}

func (b *BookCopy) IsOnLoan() bool {
	return b.availability == BookAvailabilityOnLoan
}

func (b *BookCopy) Loan() {
	b.availability = BookAvailabilityOnLoan
}

func NewBookCopy(id ID, availability BookAvailability) (*BookCopy, error) {
	return &BookCopy{
		id:           id,
		availability: availability,
	}, nil
}

type ISBN struct {
	value uint
}

func (i ISBN) String() int {
	return int(i.value)
}

func NewISBN(value int) (ISBN, error) {
	return ISBN{
		value: uint(value),
	}, nil
}

type BookFormat struct {
	value string
}

func (b BookFormat) String() string {
	return b.value
}

var (
	BookFormatHardCover = BookFormat{"hardcover"}
	BookFormatPaperBack = BookFormat{"paperback"}
	BookFormatAudioBook = BookFormat{"audiobook"}
)

func NewBookFormat(format string) (BookFormat, error) {
	switch format {
	case BookFormatHardCover.value:
		return BookFormatHardCover, nil
	case BookFormatPaperBack.value:
		return BookFormatPaperBack, nil
	case BookFormatAudioBook.value:
		return BookFormatAudioBook, nil
	default:
		return BookFormat{}, fmt.Errorf("%s is not a valid bookCopy format", format)
	}
}

type BookAvailability struct {
	value string
}

func (b BookAvailability) String() string {
	return b.value
}

var (
	BookAvailabilityAvailable = BookAvailability{"available"}
	BookAvailabilityOnLoan    = BookAvailability{"on_loan"}
	BookAvailabilityLost      = BookAvailability{"lost"}
)

func NewBookAvailability(availability string) (BookAvailability, error) {
	switch availability {
	case BookAvailabilityAvailable.value:
		return BookAvailabilityAvailable, nil
	case BookAvailabilityOnLoan.value:
		return BookAvailabilityOnLoan, nil
	case BookAvailabilityLost.value:
		return BookAvailabilityLost, nil
	default:
		return BookAvailability{}, fmt.Errorf("%s is not a valid availability", availability)
	}
}
