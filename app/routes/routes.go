package routes

import (
	"yukevent/controllers/events"
	"yukevent/controllers/organizers"
	"yukevent/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware       middleware.JWTConfig
	UserController      users.UserController
	OrganizerController organizers.OrganizerController
	EventController     events.EventController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	users := e.Group("users")
	users.POST("/register", cl.UserController.Register)
	users.POST("/login", cl.UserController.Login)

	organizers := e.Group("organizers")
	organizers.POST("/register", cl.OrganizerController.Register)
	organizers.POST("/login", cl.OrganizerController.Login)
	organizers.POST("/create-event", cl.EventController.Create, middleware.JWTWithConfig(cl.JWTMiddleware))
}
