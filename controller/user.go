package controller

import (
	"message-board/dao"
	"message-board/model"
	"message-board/service"
	"message-board/utils"
	"net/http"
)

func CtrlUserRegister(name, email, password string) (err error, resp utils.RespData) {
	var accessToken, refreshToken string
	var uid int64

	err, accessToken, refreshToken, uid = service.RegisterAccount(name, email, password)

	if err != nil {
		return err, utils.RespData{}
	}

	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data: struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
			Uid          int64  `json:"uid"`
		}{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			Uid:          uid,
		},
	}
	return
}

func CtrlUserLogin(email, password string) (err error, resp utils.RespData) {
	var accessToken, refreshToken string
	var uid int64

	err, accessToken, refreshToken, uid = service.LoginAccount(email, password)

	if err != nil {
		return err, utils.RespData{}
	}

	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data: struct {
			AccessToken  string `json:"access_token"`
			RefreshToken string `json:"refresh_token"`
			Uid          int64  `json:"uid"`
		}{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			Uid:          uid,
		},
	}
	return
}

func CtrlSendVerifyCode(email string) (err error, resp utils.RespData) {
	service.SendVerifyCode(email)
	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
	}
	return
}

func CtrlChangePassword(email, password, verify string) (err error, resp utils.RespData) {
	if err := utils.VerifyInputCode(email, verify); err != nil {
		return err, utils.RespData{}
	}
	err = service.ChangePass(email, password)
	if err != nil {
		return err, utils.RespData{}
	}
	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
	}
	return
}

func CtrlChangeDetail(user *model.User) (err error, resp utils.RespData) {
	if err = service.ChangeUserDetail(user); err != nil {
		return err, utils.RespData{}
	}
	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
	}
	return
}

func CtrlGetDetail(uid *uint, name, email *string) (err error, resp utils.RespData) {
	user := dao.User{}
	if uid != nil {
		err = service.GetUserDetailById(&user, *uid)
		if err != nil {
			return err, utils.RespData{}
		}
	} else if name != nil {
		err = service.GetUserDetailByName(&user, *name)
		if err != nil {
			return err, utils.RespData{}
		}
	} else {
		err = service.GetUserDetailByEmail(&user, *email)
		if err != nil {
			return err, utils.RespData{}
		}
	}
	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       user.User,
	}
	return
}
