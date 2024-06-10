package main

import (
	_ "embed"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/subscribeddotdev/subscribed-backend/tests/jwks"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	router := echo.New()

	router.Use(middleware.Logger())
	router.GET("/jwks", func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		return c.String(http.StatusOK, fmt.Sprintf(`{"keys": [%s]}`, jwks.JwksPublicKey))
	})

	err := router.Start(":8090")
	if err != nil {
		logger.Error("emulators crashed", err)
		os.Exit(1)
	}
}
