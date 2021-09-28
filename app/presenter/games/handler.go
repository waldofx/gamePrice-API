package games

import (
	"gameprice-api/app/presenter/games/request"
	"gameprice-api/app/presenter/games/response"
	"gameprice-api/business/games"
	"net/http"

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

func (handler *Presenter) Insert(echoContext echo.Context) error {
	var req request.GameInsert
	if err := echoContext.Bind(&req); err != nil {
		return echoContext.JSON(http.StatusBadRequest, "something wrong")
	}

	domain := request.ToDomain(req)
	resp, err := handler.serviceGame.Append(domain)
	if err != nil {
		return echoContext.JSON(http.StatusInternalServerError, "something wrong")
	}

	return echoContext.JSON(http.StatusOK, response.FromDomain(*resp))
}
