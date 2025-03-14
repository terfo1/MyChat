package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/terfo1/news/internal/controllers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", controllers.HomePage)
	api.Post("/signup", controllers.SignUp)
}
