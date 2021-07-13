package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lxxonx/golang-fiber/controllers"
)

func postRoute(app *fiber.App) {
	app.Get("/api/posts", controllers.GetPosts)
	app.Post("/api/post/create", controllers.CreatePost)
	app.Delete("/api/post/:id", controllers.DeletePost)
}