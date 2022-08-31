package model

import "time"

type Message struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Author    User      `gorm:"embedded" json:"author"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
