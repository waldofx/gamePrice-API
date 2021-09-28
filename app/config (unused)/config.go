package config

import (
	"fmt"
	"gameprice-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	db := "project"
	connectionString := fmt.Sprintf("root:@tcp(0.0.0.0:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", db)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	InitMigrate()
}

func InitMigrate() {
	//DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Game{})
}