package main

import (
	"message-board/api"
	"message-board/dao"
	"message-board/utils"
	"time"
)

func main() {

	utils.EnableLog()
	<-time.NewTimer(time.Second).C // 延迟 1s，让日志启动

	dao.InitDao()
	api.InitRouter()
}
