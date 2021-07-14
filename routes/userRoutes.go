package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lxxonx/golang-fiber/controllers"
)
func userRoute(api fiber.Router) {
	users := api.Group("/users")        // /api/v1

	users.Post("", controllers.Register)
	users.Post("/login", controllers.Login)
	users.Get("/:id", controllers.UserGet)
	users.Post("/logout", controllers.Logout)
	users.Get("/", controllers.GetUsers)
}