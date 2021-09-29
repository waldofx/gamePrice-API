package games

import (
	"gameprice-api/app/presenter/games/request"
	"gameprice-api/app/presenter/games/response"
	"gameprice-api/business/games"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceGame games.Service
}

func NewHandler(gameServ games.Service) *Presenter {
	return &Presenter{
		serviceGame: gameServ,
	}
}

func (handler *Presenter) Create(echoContext echo.Context) error {
	var req request.GameInsert
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

	return echoContext.JSON(http.StatusOK, response.FromDomain(*resp))
}

func (handler *Presenter) Update(echoContext echo.Context) error{
	var req request.GameUpdate
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Bad Request",
		})
	}
	domain := request.ToDomainUpdate(req)
	resp, err := handler.serviceGame.Update(domain)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Internal Server Error",
		})
	}
	return echoContext.JSON(http.StatusOK, response.FromDomain(*resp))
}

func (handler *Presenter) ReadAll(echoContext echo.Context) error{
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

func (handler *Presenter) ReadID(echoContext echo.Context) error {
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

func(handler *Presenter) Delete(echoContext echo.Context) error{
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
			"message": "bad request 2 boogaloo",
		})
	}
	
	return echoContext.JSON(http.StatusOK, result)
}