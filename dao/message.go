package dao

func InsertMsg(message Message) (err error) {
	err = dB.Create(message).Error
	return
}

func SelectTopMessages(message *[]Message) (err error) {
	err = dB.Where("parent_id = ?", nil).Find(message).Error
	return
}

func SelectChildrenMsgById(message *[]Message, mid uint) (err error) {
	err = dB.Where("parent_id = ?", mid).Find(message).Error
	return
}

func SelectMessageById(message *Message, mid uint) (err error) {
	err = dB.Where("id = ?", mid).First(message).Error
	return
}
