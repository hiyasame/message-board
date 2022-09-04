package model

import "time"

type Message struct {
	ID        uint      `json:"id"`
	Author    User      `json:"author"`
	Message   string    `json:"message"`
	ParentId  uint      `json:"parent_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
