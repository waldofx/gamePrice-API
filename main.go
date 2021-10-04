package main

import (
	"log"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	_handlerGames "gameprice-api/app/presenter/games"
	_servGames "gameprice-api/business/games"
	_repoGames "gameprice-api/repository/mysql/games"

	_handlerSellers "gameprice-api/app/presenter/sellers"
	_servSellers "gameprice-api/business/sellers"
	_repoSellers "gameprice-api/repository/mysql/sellers"

	_handlerUsers "gameprice-api/app/presenter/users"
	_servUsers "gameprice-api/business/users"
	_repoUsers "gameprice-api/repository/mysql/users"

	_repoGOGapis "gameprice-api/repository/thirdparty/gogapis"
	_repoSteamapis "gameprice-api/repository/thirdparty/steamapis"

	_handlerProducts "gameprice-api/app/presenter/products"
	_servProducts "gameprice-api/business/products"
	_repoProducts "gameprice-api/repository/mysql/products"

	_handlerWishes "gameprice-api/app/presenter/wishes"
	_servWishes "gameprice-api/business/wishes"
	_repoWishes "gameprice-api/repository/mysql/wishes"

	_dbDriver "gameprice-api/repository/mysql"

	_middleware "gameprice-api/app/middleware"
	_routes "gameprice-api/app/routes"
)

func init() {
	viper.SetConfigFile(`app/config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_repoGames.Games{},
		&_repoSellers.Sellers{},
		&_repoUsers.Users{},
		&_repoProducts.Products{},
		&_repoWishes.Wishes{},
	)
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}
	db := configDB.InitialDB()
	dbMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second


	e := echo.New()

	// factory of domain
	gamesRepo := _repoGames.NewRepoMySQL(db)
	gamesService := _servGames.NewService(gamesRepo)
	gamesHandler := _handlerGames.NewHandler(gamesService)
	sellersRepo := _repoSellers.NewRepoMySQL(db)
	sellersService := _servSellers.NewService(sellersRepo)
	sellersHandler := _handlerSellers.NewHandler(sellersService)
	usersRepo := _repoUsers.NewRepoMySQL(db)
	usersService := _servUsers.NewService(usersRepo, &configJWT, timeoutContext)
	usersHandler := _handlerUsers.NewHandler(usersService)
	steamapisRepo := _repoSteamapis.NewRepo()
	gogapisRepo := _repoGOGapis.NewRepo()
	productsRepo := _repoProducts.NewRepoMySQL(db)
	productsService := _servProducts.NewService(productsRepo, steamapisRepo, gogapisRepo)
	productsHandler := _handlerProducts.NewHandler(productsService)
	wishesRepo := _repoWishes.NewRepoMySQL(db)
	wishesService := _servWishes.NewService(wishesRepo)
	wishesHandler := _handlerWishes.NewHandler(wishesService)

	// initial of routes
	routesInit := _routes.HandlerList{
		JWTMiddleware:  configJWT.Init(),
		GameHandler: *gamesHandler,
		SellerHandler: *sellersHandler,
		UserHandler: *usersHandler,
		ProductHandler: *productsHandler,
		WishHandler: *wishesHandler, //wish masih kurang
	}
	routesInit.RouteRegister(e)


	_middleware.LogMiddlewareInit(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
