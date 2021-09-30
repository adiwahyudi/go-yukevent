package users

import (
	"net/http"
	"yukevent/business/users"
	"yukevent/controllers"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService users.Service
}

func NewControllerUser(serv users.Service) *UserController {
	return &UserController{
		userService: serv,
	}
}

func (ctrl *UserController) Register(c echo.Context) error {

	ctx := c.Request().Context()
	registerReq := users.Domain{}
	if err := c.Bind(&registerReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := ctrl.userService.Register(ctx, registerReq.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, result)

}
