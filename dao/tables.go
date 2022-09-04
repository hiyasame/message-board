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
	AuthorId uint   `gorm:"not null"`
	Message  string `gorm:"not null"`
}

func (m *Message) AsModel() (err error, msg model.Message) {
	var user User
	err = SelectUserById(&user, m.AuthorId)
	if err != nil {
		return
	}
	msg = model.Message{
		ID:        m.Base.ID,
		Author:    user.User,
		Message:   m.Message,
		ParentId:  m.ParentId,
		CreatedAt: m.Base.CreatedAt,
		UpdatedAt: m.Base.UpdatedAt,
	}
	return
}
