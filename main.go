package main

import (
	"net/http"

	"github.com/GrubbyRM/CyberGo/cmd/handlers"
	"github.com/GrubbyRM/CyberGo/cmd/storage"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})
	storage.InitDB()
	e.POST("/users", handlers.CreateUser)
	e.Logger.Fatal(e.Start(":1331"))
}
