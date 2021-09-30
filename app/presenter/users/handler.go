package users

import (
	"gameprice-api/app/presenter/users/request"
	"gameprice-api/app/presenter/users/response"
	"gameprice-api/business/users"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceUser users.Service
}

func NewHandler(userServ users.Service) *Presenter {
	return &Presenter{
		serviceUser: userServ,
	}
}

func (handler *Presenter) Create(echoContext echo.Context) error {
	var req request.Users
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}
	domain := request.ToDomain(req)
	resp, err := handler.serviceUser.Append(domain)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}

	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"message": "Data Inserted.",
		"data": response.FromDomain(*resp),
	})
}

func (handler *Presenter) Update(echoContext echo.Context) error{
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

func (handler *Presenter) ReadAll(echoContext echo.Context) error{
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

func (handler *Presenter) ReadID(echoContext echo.Context) error {
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

func(handler *Presenter) Delete(echoContext echo.Context) error{
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