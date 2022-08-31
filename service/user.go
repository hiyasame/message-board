package service

import (
	"golang.org/x/crypto/bcrypt"
	"message-board/dao"
	"message-board/model"
	"message-board/utils"
)

func RegisterAccount(name, email, password string) (err error, accessToken, refreshToken string) {
	salt := utils.GenerateUUIDStr()
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := dao.User{
		User: model.User{
			Name:  name,
			Email: email,
		},
		PasswordSalt:      salt,
		PasswordEncrypted: string(encrypted),
	}
	err = dao.InsertUser(&user)
	if err != nil {
		return
	}
	accessToken, refreshToken, err = utils.GenerateTokenPair(int64(user.User.ID))
	return
}
