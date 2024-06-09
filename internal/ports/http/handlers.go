package http

import (
	"github.com/labstack/echo/v4"
	"github.com/subscribeddotdev/subscribed-backend/internal/app"
)

type handlers struct {
	application *app.App
}

func (h handlers) GetHelloWorld(ctx echo.Context) error {
	//TODO implement me
	panic("implement me")
}
