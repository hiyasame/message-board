package controller

import (
	"message-board/dao"
	"message-board/model"
	"message-board/service"
	"message-board/utils"
	"net/http"
)

func CtrlMessagePut(message dao.Message) (err error, resp utils.RespData) {
	err = service.PutMessage(message)
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

func CtrlMessageGet(id uint) (err error, resp utils.RespData) {
	var message dao.Message
	err = service.GetMessage(&message, id)
	if err != nil {
		return err, utils.RespData{}
	}
	err, mod := message.AsModel()
	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       mod,
	}
	return
}

func CtrlMessagesGet() (err error, resp utils.RespData) {
	var _messages []dao.Message
	err = service.GetTopMessages(&_messages)
	if err != nil {
		return err, utils.RespData{}
	}
	messages := utils.Map(_messages, func(t dao.Message) model.Message {
		_, mod := t.AsModel()
		return mod
	})
	resp = utils.RespData{
		HttpStatus: http.StatusOK,
		Status:     20000,
		Info:       utils.InfoSuccess,
		Data:       messages,
	}
	return
}
