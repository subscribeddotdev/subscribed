package http

import (
	"github.com/labstack/echo/v4"
	"github.com/subscribeddotdev/subscribed-backend/internal/app"
)

type handlers struct {
	application *app.App
}

func (h handlers) GetHelloWorld(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"msg": "Hello World",
	})
}
