package wishes

import (
	"gameprice-api/app/presenter/wishes/request"
	"gameprice-api/app/presenter/wishes/response"
	"gameprice-api/business/wishes"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceWish wishes.Service
}

func NewHandler(wisheserv wishes.Service) *Presenter {
	return &Presenter{
		serviceWish: wisheserv,
	}
}

func (handler *Presenter) Create(echoContext echo.Context) error {
	var req request.Wishes
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}
	domain := request.ToDomain(req)
	resp, err := handler.serviceWish.Append(domain)
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
	var req request.Wishes
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}
	domain := request.ToDomain(req)
	resp, err := handler.serviceWish.Update(domain, id)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}
	return echoContext.JSON(http.StatusOK, response.FromDomain(*resp))
}

func (handler *Presenter) ReadAll(echoContext echo.Context) error{
	wishes, err := handler.serviceWish.FindAll()
	if err != nil {
		return echoContext.JSON(http.StatusBadRequest,  map[string]interface{}{
			"message": "Bad Request",
		})
	}
	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"wishes": response.NewResponseArray(wishes),
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
	resp, err :=  handler.serviceWish.FindByID(id)
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
	wish, err1 := handler.serviceWish.FindByID(id)
	result, err2 :=  handler.serviceWish.Delete(wish, id)

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