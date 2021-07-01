package database

import (
	"log"
	"os"

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
	dsn := os.Getenv("DSN")

    connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
    	panic("failed to connect database")
    }
	DB = connection
	
	connection.AutoMigrate(&models.User{})
}