package models

import (
	"time"
)

type Message struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	SenderID   uint      `gorm:"not null" json:"sender_id"`
	ReceiverID uint      `gorm:"not null" json:"receiver_id"`
	Content    string    `gorm:"type:text;not null" json:"content"`
	CreatedAt  time.Time `json:"created_at"`

	Sender   User `gorm:"foreignKey:SenderID" json:"-"`
	Receiver User `gorm:"foreignKey:ReceiverID" json:"-"`
}
