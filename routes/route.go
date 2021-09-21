package routes

import (
	"yukevent/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()

	e.POST("/register", controllers.RegisterUserController)

	return e
}
