package controllers

import (
	"net/http"
	"yukevent/lib/database"

	"github.com/labstack/echo/v4"
)

func RegisterUserController(c echo.Context) error {

	var userReq RequestUserRegister

	c.Bind(&userReq)

	result, err := database.RegisterUser(userReq.toModel())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, newResponse(*result))
}
