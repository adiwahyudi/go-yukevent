package routes

import (
	"yukevent/controllers"
	m "yukevent/middlewares"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {

	e := echo.New()

	e.POST("/register", controllers.RegisterUserController)

	m.LogMiddleware(e)
	return e
}
