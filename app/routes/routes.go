package routes

import (
	"gameprice-api/app/presenter/games"

	"github.com/labstack/echo/v4"
)

type HandlerList struct {
	GameHandler games.Presenter
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	games := e.Group("games")
	games.POST("/register", handler.GameHandler.Insert)
}
