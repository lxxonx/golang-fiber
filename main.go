package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/lxxonx/golang-fiber/database"
	"github.com/lxxonx/golang-fiber/routes"
)

func main() {
  // db connect
  database.Connect()

  // fiber create
  app := fiber.New(fiber.Config{
    BodyLimit: 100 * 1024 * 1024, // 100MB 
    //default limit of 4MB
  })  

  app.Use(logger.New())

  app.Use(cors.New(cors.Config{
	  AllowCredentials: true,
	}))
  
  // routing
  routes.Setup(app)

  // fiber start 
  app.Listen(":3000")
}