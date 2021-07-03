package database

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/lxxonx/golang-fiber/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	dsn := "postgres://lxxonx@localhost:5432/test"
	fmt.Print(dsn)

    connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
    	panic("failed to connect database")
    }
	DB = connection
	
	// migrate
	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.Post{})

}