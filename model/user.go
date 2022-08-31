package model

type User struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Name   string `gorm:"unique;not null" json:"name"`
	Avatar string `json:"avatar"`
	Bio    string `json:"bio"`
	Email  string `gorm:"unique;not null" json:"email"`
}
