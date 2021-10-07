package routes

import (
	"net/http"
	middlewareApp "yukevent/app/middleware"
	"yukevent/business"
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

	// Public
	e.GET("/events", cl.EventController.AllEvent)
	e.GET("/event/:id", cl.EventController.EventByID)
	e.GET("/:organizerID/events", cl.EventController.EventByIdOrganizer)

	// Users
	users := e.Group("users")
	users.POST("/register", cl.UserController.Register)
	users.POST("/login", cl.UserController.Login)

	// Organizers
	organizers := e.Group("organizers")
	organizers.POST("/register", cl.OrganizerController.Register)
	organizers.POST("/login", cl.OrganizerController.Login)

	organizers.POST("/create-event", cl.EventController.Create, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationOrganizer())
	organizers.PUT("/update-event/:id", cl.EventController.Update, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationOrganizer())
	organizers.DELETE("/delete-event/:id", cl.EventController.Delete, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationOrganizer())
	organizers.GET("/my-events", cl.EventController.MyEventByOrganizer, middleware.JWTWithConfig(cl.JWTMiddleware), RoleValidationOrganizer())

}

func RoleValidationUser() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			claims := middlewareApp.GetUser(c)

			if claims.Role == "user" {
				return hf(c)
			} else {
				return controller.NewErrorResponse(c, http.StatusForbidden, business.ErrUnathorized)
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
				return controller.NewErrorResponse(c, http.StatusForbidden, business.ErrUnathorized)
			}
		}
	}
}
