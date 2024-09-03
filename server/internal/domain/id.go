package domain

import (
	"fmt"
	"strings"

	"github.com/oklog/ulid/v2"
)

type ID string

func NewID() ID {
	return ID(ulid.Make().String())
}

func NewIdFromString(value string) (ID, error) {
	parsedID, err := ulid.Parse(value)
	if err != nil {
		return "", err
	}

	return ID(parsedID.String()), nil
}

func NewIdFromStringWithPrefix(value string) (ID, error) {
	parts := strings.Split(value, "_")
	parsedID, err := ulid.Parse(parts[1])
	if err != nil {
		return "", err
	}

	return ID(fmt.Sprintf("%s_%s", parts[0], parsedID.String())), nil
}

func (i ID) String() string {
	return string(i)
}

func (i ID) WithPrefix(prefix string) ID {
	id := ID(fmt.Sprintf("%s_%s", prefix, i))
	return id
}

func (i ID) IsEmpty() bool {
	return string(i) == ""
}

func (i ID) ToMessageID() MessageID {
	return MessageID(fmt.Sprintf("msg_%s", i))
}
