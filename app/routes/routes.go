package routes

import (
	"gameprice-api/controllers/games"
	"gameprice-api/controllers/products"
	"gameprice-api/controllers/sellers"
	"gameprice-api/controllers/users"
	"gameprice-api/controllers/wishes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerList struct {
	JWTMiddleware  middleware.JWTConfig
	GameHandler games.Controller
	SellerHandler sellers.Controller
	UserHandler users.Controller
	ProductHandler products.Controller
	WishHandler wishes.Controller
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
	wishes.POST("/insert", handler.WishHandler.Create, middleware.JWTWithConfig(handler.JWTMiddleware))
	wishes.GET("/all", handler.WishHandler.ReadAll, middleware.JWTWithConfig(handler.JWTMiddleware))
	wishes.GET("/user/:id", handler.WishHandler.ReadUserID, middleware.JWTWithConfig(handler.JWTMiddleware))
	wishes.GET("/:id", handler.WishHandler.ReadID, middleware.JWTWithConfig(handler.JWTMiddleware))
	wishes.PUT("/:id", handler.WishHandler.Update, middleware.JWTWithConfig(handler.JWTMiddleware))
	wishes.DELETE("/:id", handler.WishHandler.Delete, middleware.JWTWithConfig(handler.JWTMiddleware))
}
