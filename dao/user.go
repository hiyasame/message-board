package dao

import (
	"golang.org/x/crypto/bcrypt"
	"message-board/model"
	"message-board/utils"
)

func InsertUser(user *User) error {
	result := dB.Create(user)
	return result.Error
}

func SelectUserByEmail(user *User, email string) (err error) {
	result := dB.First(user, "email = ?", email)
	err = result.Error
	return
}

func SelectUserByName(user *User, name string) (err error) {
	result := dB.First("name = ?", name)
	err = result.Error
	return
}

func SelectUserById(user *User, id uint) (err error) {
	result := dB.First("id = ?", id)
	err = result.Error
	return
}

func UpdateUserPassword(email, password string) (err error) {
	salt := utils.GenerateUUIDStr()
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
	err = dB.Model(&User{}).Where("email = ?", email).Updates(User{
		PasswordSalt:      salt,
		PasswordEncrypted: string(encrypted),
	}).Error
	return
}

func UpdateUser(user *model.User) (err error) {
	err = dB.Model(&User{
		User: model.User{
			ID: user.ID,
		},
	}).Updates(&User{
		User: *user,
	}).Error
	return
}
