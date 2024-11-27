package handlers

import (
	"fmt"
	"net/http"

	"github.com/GrubbyRM/CyberGo/cmd/models"
	"github.com/GrubbyRM/CyberGo/cmd/repositores"
	"github.com/labstack/echo/v4"
)

func CreateUser(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	fmt.Println("Error occured!")
	newUser, err := repositores.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	fmt.Println("OK!")
	return c.JSON(http.StatusCreated, newUser)
}
