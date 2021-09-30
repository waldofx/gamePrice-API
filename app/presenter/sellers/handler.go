package sellers

import (
	"gameprice-api/app/presenter/sellers/request"
	"gameprice-api/app/presenter/sellers/response"
	"gameprice-api/business/sellers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceSeller sellers.Service
}

func NewHandler(sellerServ sellers.Service) *Presenter {
	return &Presenter{
		serviceSeller: sellerServ,
	}
}

func (handler *Presenter) Create(echoContext echo.Context) error {
	var req request.SellerInsert
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}
	domain := request.ToDomain(req)
	resp, err := handler.serviceSeller.Append(domain)
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
	var req request.SellerUpdate
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}
	domain := request.ToDomainUpdate(req)
	resp, err := handler.serviceSeller.Update(domain)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}
	return echoContext.JSON(http.StatusOK, response.FromDomain(*resp))
}

func (handler *Presenter) ReadAll(echoContext echo.Context) error{
	sellers, err := handler.serviceSeller.FindAll()
	if err != nil {
		return echoContext.JSON(http.StatusBadRequest,  map[string]interface{}{
			"message": "Bad Request",
		})
	}
	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"sellers": response.NewResponseArray(sellers),
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
	resp, err :=  handler.serviceSeller.FindByID(id)
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
	seller, err1 := handler.serviceSeller.FindByID(id)
	result, err2 :=  handler.serviceSeller.Delete(seller, id)

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