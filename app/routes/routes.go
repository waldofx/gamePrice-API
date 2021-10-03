package routes

import (
	"gameprice-api/app/presenter/games"
	"gameprice-api/app/presenter/products"
	"gameprice-api/app/presenter/sellers"
	"gameprice-api/app/presenter/users"
	"gameprice-api/app/presenter/wishes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerList struct {
	JWTMiddleware  middleware.JWTConfig
	GameHandler games.Presenter
	SellerHandler sellers.Presenter
	UserHandler users.Presenter
	ProductHandler products.Presenter
	WishHandler wishes.Presenter
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
	users.POST("/register", handler.UserHandler.Store)
	users.GET("/token", handler.UserHandler.CreateToken)
	// users.POST("/insert", handler.UserHandler.Create)
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

	wishes := e.Group("wishes")
	wishes.POST("/insert", handler.WishHandler.Create)
	wishes.GET("/all", handler.WishHandler.ReadAll)
	wishes.GET("/:id", handler.WishHandler.ReadID)
	wishes.PUT("/:id", handler.WishHandler.Update)
	wishes.DELETE("/:id", handler.WishHandler.Delete)
}
