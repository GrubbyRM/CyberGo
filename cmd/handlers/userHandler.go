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
	newUser, err := repositores.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	fmt.Println("OK!")
	return c.JSON(http.StatusCreated, newUser)
}

// TODO: add endpoint
func DeleteUser(c echo.Context) error {
	userName := c.Param("name")

	err := repositores.DeleteUser(userName)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusNoContent)
}

// TODO: add endpoint
func ShowUsers(c echo.Context) error {
	users, err := repositores.ShowUsers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, users)
}
