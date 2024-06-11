package iam

import (
	"net/mail"
	"strings"
)

type Email struct {
	address string
}

func NewEmail(address string) (Email, error) {
	_, err := mail.ParseAddress(address)
	if err != nil {
		return Email{}, err
	}

	return Email{
		address: strings.TrimSpace(address),
	}, nil
}

func (e Email) String() string {
	return e.address
}

func (e Email) IsEmpty() bool {
	return e.address == ""
}
