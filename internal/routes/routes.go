package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/terfo1/news/internal/transport/http"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", http.HomePage)
	api.Get("/register", http.Register)
}
