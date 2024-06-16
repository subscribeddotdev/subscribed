package domain

import (
	"strings"
	"time"

	"errors"
)

type Environment struct {
	id         ID
	orgID      ID
	name       string
	envType    EnvType
	createdAt  time.Time
	archivedAt *time.Time
}

func NewEnvironment(name string, orgID ID, envType EnvType) (*Environment, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	err := envType.Validate()
	if err != nil {
		return nil, err
	}

	if orgID.IsEmpty() {
		return nil, errors.New("orgID cannot be empty")
	}

	return &Environment{
		id:         NewID(),
		orgID:      orgID,
		name:       name,
		envType:    envType,
		createdAt:  time.Now(),
		archivedAt: nil,
	}, nil
}

func (e *Environment) Id() ID {
	return e.id
}

func (e *Environment) Name() string {
	return e.name
}

func (e *Environment) OrgID() ID {
	return e.orgID
}

func (e *Environment) CreatedAt() time.Time {
	return e.createdAt
}

func (e *Environment) ArchivedAt() *time.Time {
	return e.archivedAt
}

func (e *Environment) Type() EnvType {
	return e.envType
}

type EnvType struct {
	name string
}

func (e EnvType) String() string {
	return e.name
}

var (
	EnvTypeProduction  = EnvType{"production"}
	EnvTypeDevelopment = EnvType{"development"}
)

func (e EnvType) Validate() error {
	_, err := marshallToEventType(e.name)
	return err
}

func marshallToEventType(name string) (EnvType, error) {
	name = strings.TrimSpace(name)
	if name == "" {
		return EnvType{}, errors.New("environment type cannot be empty")
	}

	switch name {
	case EnvTypeProduction.name:
		return EnvTypeProduction, nil
	case EnvTypeDevelopment.name:
		return EnvTypeDevelopment, nil
	default:
		return EnvType{}, errors.New("invalid environment type")
	}
}
