package domain

import "context"

type EnvironmentRepository interface {
	Insert(ctx context.Context, env *Environment) error
}
