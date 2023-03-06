package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"hello": "world",
		})
	})

	e.GET("/v2/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"hello": "world",
		})
	})

	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
