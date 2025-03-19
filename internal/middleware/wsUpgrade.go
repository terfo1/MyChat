package middleware

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/terfo1/news/internal/models"
	"log"
)

func WsUpgrade(c *fiber.Ctx) error {
	if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		user := c.Locals("user").(models.User)
		//c.Locals("GROUP", string(c.Request().Header.Peek("GROUP")))
		//c.Locals("USER", string(c.Request().Header.Peek("USER")))
		log.Print("WebSocket upgrade successful for user:", user.Email)
		return c.Next()
	}
	return fiber.ErrUpgradeRequired
}
