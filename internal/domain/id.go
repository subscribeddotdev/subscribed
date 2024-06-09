package domain

import "github.com/oklog/ulid/v2"

type ID struct {
	value ulid.ULID
}

func NewID() ID {
	return ID{
		value: ulid.Make(),
	}
}

func NewIdFromString(id string) (ID, error) {
	value, err := ulid.Parse(id)
	if err != nil {
		return ID{}, err
	}

	return ID{
		value: value,
	}, nil
}

func (i ID) String() string {
	return i.value.String()
}
