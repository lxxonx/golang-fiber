package routes

import (
	"github.com/gofiber/fiber/v2"
)


func Setup(app *fiber.App) {
	api := app.Group("/api")      // /api


	// Static file server
	app.Static("/", "./files")
	// => http://localhost:3000/hello.txt
	// => http://localhost:3000/gopher.gif
	
	userRoute(api)
	postRoute(api)
}


