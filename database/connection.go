package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lxxonx/golang-fiber/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
	dsn := "postgres://lxxonx@localhost:5432/test"
	fmt.Print(dsn)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
		  SlowThreshold:              time.Second,   // Slow SQL threshold
		  LogLevel:                   logger.Info, // Log level
		  IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
		  Colorful:                  true,          // Disable color
		},
	  )
    connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,

	})
    if err != nil {
    	panic("failed to connect database")
    }
	DB = connection
	
	// migrate
	connection.AutoMigrate(&models.User{}, &models.Post{})

}