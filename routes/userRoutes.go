package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lxxonx/golang-fiber/controllers"
)
func userRoute(api fiber.Router) {
	users := api.Group("/users")        // /api/users

	users.Get("/", controllers.GetUsers)
	users.Get("/:id", controllers.GetUser)
	users.Post("/", controllers.CreateUser)
	users.Post("/login", controllers.Login)
	users.Post("/logout", controllers.Logout)
}