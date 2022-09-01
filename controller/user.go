package controller

import (
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
	err = service.ChangePass(email, verify)
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
