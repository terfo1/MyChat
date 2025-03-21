package database

import "github.com/terfo1/news/internal/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Message{})
}
