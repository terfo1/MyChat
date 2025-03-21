package controllers

import (
	"github.com/terfo1/news/internal/database"
	"github.com/terfo1/news/internal/models"

	"github.com/gofiber/fiber/v2"
)

func GetChatHistory(c *fiber.Ctx) error {
	user1 := c.Params("user1")
	user2 := c.Params("user2")

	var messages []models.Message
	database.DB.Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)", user1, user2, user2, user1).
		Order("created_at ASC").Find(&messages)

	return c.JSON(messages)
}
