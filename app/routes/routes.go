package routes

import (
	"gameprice-api/app/presenter/games"
	"gameprice-api/app/presenter/products"
	"gameprice-api/app/presenter/sellers"
	"gameprice-api/app/presenter/users"

	"github.com/labstack/echo/v4"
)

type HandlerList struct {
	GameHandler games.Presenter
	SellerHandler sellers.Presenter
	UserHandler users.Presenter
	ProductHandler products.Presenter
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	games := e.Group("games")
	games.POST("/insert", handler.GameHandler.Create)
	games.GET("/all", handler.GameHandler.ReadAll)
	games.GET("/:id", handler.GameHandler.ReadID)
	games.PUT("/:id", handler.GameHandler.Update)
	games.DELETE("/:id", handler.GameHandler.Delete)

	sellers := e.Group("sellers")
	sellers.POST("/insert", handler.SellerHandler.Create)
	sellers.GET("/all", handler.SellerHandler.ReadAll)
	sellers.GET("/:id", handler.SellerHandler.ReadID)
	sellers.PUT("/:id", handler.SellerHandler.Update)
	sellers.DELETE("/:id", handler.SellerHandler.Delete)

	users := e.Group("users")
	users.POST("/insert", handler.UserHandler.Create)
	users.GET("/all", handler.UserHandler.ReadAll)
	users.GET("/:id", handler.UserHandler.ReadID)
	users.PUT("/:id", handler.UserHandler.Update)
	users.DELETE("/:id", handler.UserHandler.Delete)

	products := e.Group("products")
	products.POST("/insert", handler.ProductHandler.Create)
	products.GET("/all", handler.ProductHandler.ReadAll)
	products.GET("/:id", handler.ProductHandler.ReadID)
	products.PUT("/:id", handler.ProductHandler.Update)
	products.DELETE("/:id", handler.ProductHandler.Delete)
}
