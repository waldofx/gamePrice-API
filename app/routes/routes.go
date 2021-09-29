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
	games.POST("/register", handler.GameHandler.Create)
	games.GET("/all", handler.GameHandler.ReadAll)
	games.GET("/:id", handler.GameHandler.ReadID)
	games.PUT("/update", handler.GameHandler.Update) //mysql.go Error 1146: Table 'project.domains' doesn't exist
	games.DELETE("/:id", handler.GameHandler.Delete) //mysql.go Error 1146: Table 'project.domains' doesn't exist
}
