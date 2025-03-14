package app

import (
	"github.com/terfo1/news/internal/configs"
	"github.com/terfo1/news/internal/database"
	"github.com/terfo1/news/internal/routes"
	"github.com/terfo1/news/internal/templates"
	"log"

	"github.com/gofiber/fiber/v2"
)

func NewApp() *fiber.App {
	templates.InitEngine()
	configs.LoadEnv()
	database.ConnectToDB()
	database.SyncDatabase()
	app := fiber.New(configs.NewFiberConfig())
	routes.SetupRoutes(app)
	return app
}

func Run() {
	config := configs.LoadConfig("app.yaml")
	app := NewApp()
	log.Fatal(app.Listen(":" + config.Port))
}
