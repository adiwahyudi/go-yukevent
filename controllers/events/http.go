package events

import (
	"net/http"
	"strconv"
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

func (ctrl *EventController) Update(c echo.Context) error {

	updateReq := request.Events{}

	if err := c.Bind(&updateReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	id, _ := strconv.Atoi(c.Param("id"))
	jwtGetID := middleware.GetUser(c)

	result, err := ctrl.eventService.Update(jwtGetID.ID, id, updateReq.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainUpdateEvent(result))

}

func (ctrl *EventController) Delete(c echo.Context) error {

	orgzID := middleware.GetUser(c)
	deletedId, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.eventService.Delete(orgzID.ID, deletedId)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, result)

}
func (ctrl *EventController) MyEventByOrganizer(c echo.Context) error {
	orgzID := middleware.GetUser(c)

	result, err := ctrl.eventService.MyEventByOrganizer(orgzID.ID)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromEventListDomain(result))
}

func (ctrl *EventController) EventByID(c echo.Context) error {

	eventID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.eventService.EventByID(eventID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomainAllEvent(result))
}

func (ctrl *EventController) EventByIdOrganizer(c echo.Context) error {

	orgzID, _ := strconv.Atoi(c.Param("organizerID"))

	result, err := ctrl.eventService.EventByIdOrganizer(orgzID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromEventListDomain(result))
}
