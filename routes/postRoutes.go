package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lxxonx/golang-fiber/controllers"
)

func postRoute(api fiber.Router) {
	posts:=api.Group("/posts") // /api/posts

	posts.Get("/", controllers.GetPosts)
	posts.Post("/", controllers.CreatePost)
	posts.Post("/music", controllers.UploadMusic)
	posts.Delete("/:id", controllers.DeletePost)
}