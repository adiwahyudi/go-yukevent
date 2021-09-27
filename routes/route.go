package routes

import (
	"yukevent/constants"
	"yukevent/controllers"
	m "yukevent/middlewares"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {

	e := echo.New()

	e.POST("/register", controllers.RegisterUserController)
	e.POST("/login", controllers.LoginUserController)

	// Implement AUTH using JWT
	jwt := middleware.JWT([]byte(constants.SECRECT_JWT))
	e.GET("/user", controllers.GetUser, jwt)

	m.LogMiddleware(e)
	return e
}
