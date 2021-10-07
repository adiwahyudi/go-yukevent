package transactions

import (
	"net/http"
	"strconv"
	"yukevent/app/middleware"
	"yukevent/business/transactions"
	"yukevent/controllers"
	"yukevent/controllers/transactions/request"
	"yukevent/controllers/transactions/response"

	"github.com/labstack/echo/v4"
)

type TransController struct {
	transService transactions.Service
}

func NewControllerEvent(serv transactions.Service) *TransController {
	return &TransController{
		transService: serv,
	}
}

func (ctrl *TransController) Create(c echo.Context) error {

	createReq := request.Transactions{}

	if err := c.Bind(&createReq); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	jwtGetID := middleware.GetUser(c)

	result, err := ctrl.transService.Trans(jwtGetID.ID, createReq.ToDomain())

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainCreate(result))

}

func (ctrl *TransController) GetAllTrans(c echo.Context) error {

	result, err := ctrl.transService.GetAllTrans()

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromTransListDomain(result))

}

func (ctrl *TransController) GetTransByID(c echo.Context) error {

	transID, _ := strconv.Atoi(c.Param("id"))

	result, err := ctrl.transService.GetTransByID(transID)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromDomainAllTrans(result))
}

func (ctrl *TransController) GetAllUserTrans(c echo.Context) error {

	jwtGetID := middleware.GetUser(c)

	result, err := ctrl.transService.GetAllUserTrans(jwtGetID.ID)

	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controllers.NewSuccessResponse(c, response.FromTransListDomain(result))
}
