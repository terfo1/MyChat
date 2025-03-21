package controllers

import (
	"encoding/json"
	"github.com/gofiber/contrib/websocket"
	"github.com/terfo1/news/internal/database"
	"github.com/terfo1/news/internal/models"
	"log"
	"strconv"
)

var clients = make(map[*websocket.Conn]uint)

func ChatHandler(c *websocket.Conn) {
	defer func() {
		delete(clients, c)
		c.Close()
	}()

	userID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		log.Println("Invalid user ID")
		return
	}
	clients[c] = uint(userID)
	log.Println("User connected:", userID)

	for {
		_, msg, err := c.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		var message models.Message
		log.Println("Received raw message:", string(msg))
		err = json.Unmarshal(msg, &message)
		if err != nil {
			log.Println("JSON error:", err)
			continue
		}

		var sender, receiver models.User
		if err := database.DB.First(&sender, message.SenderID).Error; err != nil {
			log.Println("Sender not found")
			continue
		}
		if err := database.DB.First(&receiver, message.ReceiverID).Error; err != nil {
			log.Println("Receiver not found")
			continue
		}

		database.DB.Create(&message)

		for conn, user := range clients {
			if user == message.ReceiverID {
				err = conn.WriteJSON(message)
				if err != nil {
					log.Println("Write error:", err)
				}
			}
		}
	}
}
