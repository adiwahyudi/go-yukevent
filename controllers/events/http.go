package events

import (
	"net/http"
	"yukevent/app/middleware"
	"yukevent/business/events"
	"yukevent/controllers"
	"yukevent/controllers/events/request"
	"yukevent/controllers/events/response"

	"github.com/labstack/echo/v4"
)

type EventController struct {
	eventService events.Service
}

func NewControllerEvent(serv events.Service) *EventController {
	return &EventController{
		eventService: serv,
	}
}

func (ctrl *EventController) Create(c echo.Context) error {

	createReq := request.Events{}

	if err := c.Bind(&createReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	jwtGetID := middleware.GetUser(c)

	result, err := ctrl.eventService.Create(jwtGetID.ID, createReq.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainCreate(result))

}

func (ctrl *EventController) AllEvent(c echo.Context) error {

	result, err := ctrl.eventService.AllEvent()

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromEventListDomain(result))

}
