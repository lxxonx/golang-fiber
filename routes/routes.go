package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lxxonx/golang-fiber/controllers"
)


func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user/:id", controllers.User)
	app.Post("/api/logout", controllers.Logout)
}