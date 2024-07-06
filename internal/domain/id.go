package domain

import (
	"fmt"

	"github.com/oklog/ulid/v2"
)

type ID string

func NewID() ID {
	return ID(ulid.Make().String())
}

func (i ID) WithPrefix(prefix string) ID {
	id := ID(fmt.Sprintf("%s_%s", prefix, i))
	return id
}

func NewIdFromString(value string) (ID, error) {
	parsedID, err := ulid.Parse(value)
	if err != nil {
		return "", err
	}

	return ID(parsedID.String()), nil
}

func (i ID) String() string {
	return string(i)
}

func (i ID) IsEmpty() bool {
	return string(i) == ""
}
