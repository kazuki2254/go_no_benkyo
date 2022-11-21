package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func handle(c echo.Context) error {

	return c.String(http.StatusOK, "hello world")
}
func main() {
	e := echo.New()
	e.GET("/", handle)
	e.Logger.Fatal(e.Start(":8000"))
}
