package domain

import (
	"fmt"

	"github.com/oklog/ulid/v2"
)

type ID string

func NewID() ID {
	return ID(ulid.Make().String())
}

func NewIDFromString(value string) (ID, error) {
	if value == "" {
		return "", fmt.Errorf("id cannot be empty")
	}

	return ID(value), nil
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
