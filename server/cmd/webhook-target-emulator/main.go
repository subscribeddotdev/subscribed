package main

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	lock := sync.Mutex{}
	receivedEvents := map[string]map[string][]string{}

	e.POST("/webhooks/:id/:path", func(c echo.Context) error {
		id := c.Param("id")
		path := c.Param("path")

		var event string
		if err := c.Bind(&event); err != nil {
			return err
		}

		lock.Lock()
		if _, ok := receivedEvents[id]; !ok {
			receivedEvents[id] = map[string][]string{}
		}
		receivedEvents[id][path] = append(receivedEvents[id][path], event)
		defer lock.Unlock()

		return c.NoContent(200)
	})

	e.GET("/webhooks/:id/:path", func(c echo.Context) error {
		id := c.Param("id")
		path := c.Param("path")

		lock.Lock()
		defer lock.Unlock()

		if _, ok := receivedEvents[id]; !ok {
			return c.JSON(200, []string{})
		}

		if _, ok := receivedEvents[id][path]; !ok {
			return c.JSON(200, []string{})
		}

		return c.JSON(200, receivedEvents[id][path])
	})

	e.Logger.Fatal(e.Start(":8080"))
}
