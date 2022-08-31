package dao

import (
	"gorm.io/gorm"
	"message-board/model"
	"time"
)

type User struct {
	User              model.User `gorm:"embedded"`
	PasswordSalt      string     `gorm:"not null"`
	PasswordEncrypted string     `gorm:"unique;not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Message struct {
	Base     gorm.Model `gorm:"embedded"`
	ParentId uint
	AuthorId uint `gorm:"not null"`
	Message  string
}
