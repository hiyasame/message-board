package service

import "message-board/dao"

func PutMessage(message dao.Message) (err error) {
	err = dao.InsertMsg(message)
	return
}

func GetMessage(message *dao.Message, mid uint) (err error) {
	err = dao.SelectMessageById(message, mid)
	return
}

func GetTopMessages(message *[]dao.Message) (err error) {
	err = dao.SelectTopMessages(message)
	return
}

func GetChildrenMessages(message *[]dao.Message, mid uint) (err error) {
	err = dao.SelectChildrenMsgById(message, mid)
	return
}
