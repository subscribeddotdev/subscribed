package domain

import (
	"github.com/nyaruka/phonenumbers"
)

type Member struct {
	id          ID
	firstName   string
	lastName    string
	email       Email
	phoneNumber PhoneNumber
	address     Address
}

func NewMember(
	id ID,
	firstName string,
	lastName string,
	email Email,
	phoneNumber PhoneNumber,
	address Address,
) (*Member, error) {
	return &Member{
		id:          id,
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		phoneNumber: phoneNumber,
		address:     address,
	}, nil
}

func (m *Member) ID() ID {
	return m.id
}

func (m *Member) FirstName() string {
	return m.firstName
}

func (m *Member) LastName() string {
	return m.lastName
}

func (m *Member) Email() Email {
	return m.email
}

func (m *Member) PhoneNumber() PhoneNumber {
	return m.phoneNumber
}

func (m *Member) Address() Address {
	return m.address
}

type Email struct {
	value string
}

func (e Email) String() string {
	return e.value
}

func NewEmail(email string) (Email, error) {
	return Email{
		value: email,
	}, nil
}

type PhoneNumber struct {
	value *phonenumbers.PhoneNumber
}

func NewPhoneNumber(value string) (PhoneNumber, error) {
	pn, err := phonenumbers.Parse(value, "")
	if err != nil {
		return PhoneNumber{}, err
	}

	return PhoneNumber{
		value: pn,
	}, nil
}

type Address struct {
	line1    string
	line2    string
	postCode string
	country  string
	city     string
}

func NewAddress(
	line1 string,
	line2 string,
	postCode string,
	country string,
	city string,
) (Address, error) {
	return Address{
		line1:    line1,
		line2:    line2,
		postCode: postCode,
		country:  country,
		city:     city,
	}, nil
}
