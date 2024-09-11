package main

import (
	"io"
	"net/http"
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	g := e.Group("/webhook-target")

	mountWebhookTarget(g)

	e.Logger.Fatal(e.Start(":8080"))
}

func mountWebhookTarget(g *echo.Group) {
	lock := sync.Mutex{}
	receivedEvents := map[string]map[string][]string{}

	g.POST("/webhooks/:id/:path", func(c echo.Context) error {
		id := c.Param("id")
		path := c.Param("path")

		event, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		lock.Lock()
		if _, ok := receivedEvents[id]; !ok {
			receivedEvents[id] = map[string][]string{}
		}
		receivedEvents[id][path] = append(receivedEvents[id][path], string(event))
		defer lock.Unlock()

		return c.NoContent(http.StatusOK)
	})

	g.GET("/webhooks/:id/:path", func(c echo.Context) error {
		id := c.Param("id")
		path := c.Param("path")

		lock.Lock()
		defer lock.Unlock()

		if _, ok := receivedEvents[id]; !ok {
			return c.JSON(http.StatusOK, []string{})
		}

		if _, ok := receivedEvents[id][path]; !ok {
			return c.JSON(http.StatusOK, []string{})
		}

		return c.JSON(http.StatusOK, receivedEvents[id][path])
	})

	g.GET("/webhooks", func(c echo.Context) error {
		lock.Lock()
		defer lock.Unlock()

		return c.JSON(http.StatusOK, receivedEvents)
	})
}
