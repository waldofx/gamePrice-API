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

	// initial of routes
	routesInit := _routes.HandlerList{
		GameHandler: *gamesHandler,
	}
	routesInit.RouteRegister(e)


	_middlewares.LogMiddlewareInit(e)
	log.Fatal(e.Start("localhost:8080"))
}
