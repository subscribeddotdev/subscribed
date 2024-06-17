package domain

import "context"

type EnvironmentRepository interface {
	Insert(ctx context.Context, env *Environment) error
}

type ApplicationRepository interface {
	Insert(ctx context.Context, application *Application) error
}
