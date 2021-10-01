package main

import (
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
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

	_repoSteamapis "gameprice-api/repository/thirdparty/steamapis"

	_handlerProducts "gameprice-api/app/presenter/products"
	_servProducts "gameprice-api/business/products"
	_repoProducts "gameprice-api/repository/mysql/products"

	_handlerWishes "gameprice-api/app/presenter/wishes"
	_servWishes "gameprice-api/business/wishes"
	_repoWishes "gameprice-api/repository/mysql/wishes"

	_middlewares "gameprice-api/app/middlewares"
	_routes "gameprice-api/app/routes"
)

const JWT_SECRET string = "testmvc"
const JWT_EXP int = 1

func InitDB(status string) *gorm.DB {
	db := "project"
	if status == "testing" {
		db = "project-test"
	}
	connectionString := fmt.Sprintf("root:@tcp(0.0.0.0:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", db)

	var err error
	DB, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	DB.AutoMigrate(
		&_repoGames.Games{},
		&_repoSellers.Sellers{},
		&_repoUsers.Users{},
		&_repoProducts.Products{},
		&_repoWishes.Wishes{},
		&_repoSteamapis.SteamName{},
		&_repoSteamapis.SteamAPI{}, //gagal migrate disini
	)

	return DB
}

func main() {
	db := InitDB("")
	e := echo.New()

	// factory of domain
	gamesRepo := _repoGames.NewRepoMySQL(db)
	gamesService := _servGames.NewService(gamesRepo)
	gamesHandler := _handlerGames.NewHandler(gamesService)
	sellersRepo := _repoSellers.NewRepoMySQL(db)
	sellersService := _servSellers.NewService(sellersRepo)
	sellersHandler := _handlerSellers.NewHandler(sellersService)
	usersRepo := _repoUsers.NewRepoMySQL(db)
	usersService := _servUsers.NewService(usersRepo)
	usersHandler := _handlerUsers.NewHandler(usersService)
	steamapisRepo := _repoSteamapis.NewRepo()
	productsRepo := _repoProducts.NewRepoMySQL(db)
	productsService := _servProducts.NewService(productsRepo, steamapisRepo)
	productsHandler := _handlerProducts.NewHandler(productsService)
	wishesRepo := _repoWishes.NewRepoMySQL(db)
	wishesService := _servWishes.NewService(wishesRepo)
	wishesHandler := _handlerWishes.NewHandler(wishesService)

	// initial of routes
	routesInit := _routes.HandlerList{
		GameHandler: *gamesHandler,
		SellerHandler: *sellersHandler,
		UserHandler: *usersHandler,
		ProductHandler: *productsHandler,
		WishHandler: *wishesHandler,
	}
	routesInit.RouteRegister(e)


	_middlewares.LogMiddlewareInit(e)
	log.Fatal(e.Start("localhost:8080"))
}
