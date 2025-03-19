package routes

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/terfo1/news/internal/controllers"
	"github.com/terfo1/news/internal/middleware"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/", controllers.HomePage)
	api.Post("/signup", controllers.SignUp)
	api.Post("/login", controllers.Login)
	api.Get("/profile", middleware.AuthRequire, controllers.Profile)
	api.Use("/ws", middleware.AuthRequire, middleware.WsUpgrade)
	api.Get("/ws/:id", websocket.New(controllers.GetMessage))
}
