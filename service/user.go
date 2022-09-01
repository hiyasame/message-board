package service

import (
	"golang.org/x/crypto/bcrypt"
	"message-board/dao"
	"message-board/model"
	"message-board/utils"
	"net/http"
)

func RegisterAccount(name, email, password string) (err error, accessToken, refreshToken string, uid int64) {
	salt := utils.GenerateUUIDStr()
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password+salt), bcrypt.DefaultCost)
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
	uid = int64(user.User.ID)
	accessToken, refreshToken, err = utils.GenerateTokenPair(uid)
	return
}

func LoginAccount(email, password string) (err error, accessToken, refreshToken string, uid int64) {
	var user dao.User
	err = dao.SelectUserByEmail(&user, email)
	encrypted, err := bcrypt.GenerateFromPassword([]byte(password+user.PasswordSalt), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	if err = bcrypt.CompareHashAndPassword(encrypted, []byte(user.PasswordEncrypted)); err != nil {
		return utils.ServerError{
			HttpStatus: http.StatusBadRequest,
			Status:     40007,
			Info:       "invalid request",
			Detail:     "密码错误",
		}, "", "", -1
	} else {
		uid = int64(user.User.ID)
		accessToken, refreshToken, err = utils.GenerateTokenPair(uid)
	}
	return
}

func ChangePass(email, password string) (err error) {
	err = dao.UpdateUserPassword(email, password)
	return
}

func SendVerifyCode(email string) {
	utils.SendRandomVerifyCode(email)
}

func ChangeUserDetail(user model.User) {

}
