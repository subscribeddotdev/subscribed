package domain

import "context"

type EnvironmentRepository interface {
	Insert(ctx context.Context, env *Environment) error
}

type ApplicationRepository interface {
	Insert(ctx context.Context, application *Application) error
}

type EndpointRepository interface {
	Insert(ctx context.Context, endpoint *Endpoint) error
}

type MessageRepository interface {
	Insert(ctx context.Context, message *Message) error
}
