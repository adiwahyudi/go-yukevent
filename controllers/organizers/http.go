package organizers

import (
	"net/http"
	"yukevent/business/organizers"
	"yukevent/controllers"
	"yukevent/controllers/organizers/request"
	"yukevent/controllers/organizers/response"

	"github.com/labstack/echo/v4"
)

type OrganizerController struct {
	organizerService organizers.Service
}

func NewControllerOrganizer(serv organizers.Service) *OrganizerController {
	return &OrganizerController{
		organizerService: serv,
	}
}

func (ctrl *OrganizerController) Register(c echo.Context) error {

	registerReq := request.Organizer{}

	if err := c.Bind(&registerReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	result, err := ctrl.organizerService.Register(registerReq.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainRegister(result))

}

func (ctrl *OrganizerController) Login(c echo.Context) error {

	loginReq := request.OrganzierLogin{}

	if err := c.Bind(&loginReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	result, err := ctrl.organizerService.Login(loginReq.Email, loginReq.Password)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainLogin(result))
}
