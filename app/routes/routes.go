package routes

import (
	"errors"
	"net/http"
	middlewareApp "yukevent/app/middleware"
	controller "yukevent/controllers"
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
	organizers.POST("/create-event", cl.EventController.Create, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationOrganizer())
}

func RoleValidationUser() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			if claims.Role == "user" {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, errors.New("unathorized"))
			}
		}
	}
}

func RoleValidationOrganizer() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			if claims.Role == "organizer" {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, errors.New("unathorized"))
			}
		}
	}
}
