package domain

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Route(e *echo.Echo) {
	// e := echo.New()
	// e.Start(":1323")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Secure())
	e.Use(middleware.BodyLimit("2M"))

	e.GET("/heartbeat", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
