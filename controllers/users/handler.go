package users

import (
	"gameprice-api/business/users"
	controller "gameprice-api/controllers"
	"gameprice-api/controllers/users/request"
	"gameprice-api/controllers/users/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	serviceUser users.Service
}

func NewHandler(userServ users.Service) *Controller {
	return &Controller{
		serviceUser: userServ,
	}
}

func (handler *Controller) CreateToken(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()

	username := echoContext.QueryParam("username")
	password := echoContext.QueryParam("password")

	token, err := handler.serviceUser.CreateToken(ctx, username, password)
	if err != nil {
		return controller.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	response := struct {
		Token string `json:"token"`
	}{Token: token}

	return controller.NewSuccessResponse(echoContext, response)
}

func (handler *Controller) Store(echoContext echo.Context) error {
	ctx := echoContext.Request().Context()

	req := request.Users{}
	if err := echoContext.Bind(&req); err != nil {
		return controller.NewErrorResponse(echoContext, http.StatusBadRequest, err)
	}

	err := handler.serviceUser.Store(ctx, request.ToDomain(req))
	if err != nil {
		return controller.NewErrorResponse(echoContext, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(echoContext, "Successfully inserted")
}

func (handler *Controller) Update(echoContext echo.Context) error{
	idstr := echoContext.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil{
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}
	var req request.Users
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}
	domain := request.ToDomain(req)
	resp, err := handler.serviceUser.Update(domain, id)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}
	return echoContext.JSON(http.StatusOK, response.FromDomain(*resp))
}

func (handler *Controller) ReadAll(echoContext echo.Context) error{
	users, err := handler.serviceUser.FindAll()
	if err != nil {
		return echoContext.JSON(http.StatusBadRequest,  map[string]interface{}{
			"message": "Bad Request",
		})
	}
	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"users": response.NewResponseArray(users),
	})
}

func (handler *Controller) ReadID(echoContext echo.Context) error {
	idstr := echoContext.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil{
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}
	resp, err :=  handler.serviceUser.FindByID(id)
	if err != nil{
		return echoContext.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Not Found",
		})
	}
	return echoContext.JSON(http.StatusOK, response.FromDomain(*resp))
}

func(handler *Controller) Delete(echoContext echo.Context) error{
	idstr := echoContext.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil{
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}
	user, err1 := handler.serviceUser.FindByID(id)
	result, err2 :=  handler.serviceUser.Delete(user, id)

	if err1 != nil{
		return echoContext.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "Not Found",
		})
	} else if err2 != nil{
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}
	
	return echoContext.JSON(http.StatusOK, result)
}