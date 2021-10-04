package games

import (
	"gameprice-api/business/games"
	"gameprice-api/controllers/games/request"
	"gameprice-api/controllers/games/response"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	serviceGame games.Service
}

func NewHandler(gameServ games.Service) *Controller {
	return &Controller{
		serviceGame: gameServ,
	}
}

func (handler *Controller) Create(echoContext echo.Context) error {
	var req request.Games
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}
	domain := request.ToDomain(req)
	resp, err := handler.serviceGame.Append(domain)
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

func (handler *Controller) Update(echoContext echo.Context) error{
	idstr := echoContext.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil{
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}
	var req request.Games
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}
	domain := request.ToDomain(req)
	resp, err := handler.serviceGame.Update(domain, id)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}
	return echoContext.JSON(http.StatusOK, response.FromDomain(*resp))
}

func (handler *Controller) ReadAll(echoContext echo.Context) error{
	games, err := handler.serviceGame.FindAll()
	if err != nil {
		return echoContext.JSON(http.StatusBadRequest,  map[string]interface{}{
			"message": "Bad Request",
		})
	}
	return echoContext.JSON(http.StatusOK, map[string]interface{}{
		"games": response.NewResponseArray(games),
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
	resp, err :=  handler.serviceGame.FindByID(id)
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
	game, err1 := handler.serviceGame.FindByID(id)
	result, err2 :=  handler.serviceGame.Delete(game, id)

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