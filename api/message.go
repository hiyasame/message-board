package api

import (
	"github.com/gin-gonic/gin"
	"message-board/controller"
	"message-board/dao"
	"message-board/utils"
	"strconv"
)

func HandlePutMessage(ctx *gin.Context) {
	message, ok := ctx.GetPostForm("message")
	if !ok {
		utils.RespWithParamError(ctx, "message 不可为空")
		return
	}
	id, exist := ctx.Get("uid")
	if !exist {
		utils.RespWithParamError(ctx, "未登录")
		return
	}
	uid := id.(int64)
	err, resp := controller.CtrlMessagePut(dao.Message{
		AuthorId: uint(uid),
		Message:  message,
	})
	utils.Resp(ctx, err, resp)
}

func HandlePutChildMessage(ctx *gin.Context) {
	message, ok := ctx.GetPostForm("message")
	if !ok {
		utils.RespWithParamError(ctx, "message 不可为空")
		return
	}
	_uid, exist := ctx.Get("uid")
	if !exist {
		utils.RespWithParamError(ctx, "未登录")
		return
	}
	uid := uint(_uid.(int64))
	_pid := ctx.Param("id")
	pid, err := strconv.Atoi(_pid)
	if err != nil {
		utils.RespWithParamError(ctx, "id 只能为数字")
		return
	}
	err, resp := controller.CtrlMessagePut(dao.Message{
		AuthorId: uid,
		ParentId: uint(pid),
		Message:  message,
	})
	utils.Resp(ctx, err, resp)
}

func HandleGetMessage(ctx *gin.Context) {
	err, resp := controller.CtrlMessagesGet()
	utils.Resp(ctx, err, resp)
}

func HandleGetChildrenMessage(ctx *gin.Context) {
	_pid := ctx.Param("id")
	pid, err := strconv.Atoi(_pid)
	if err != nil {
		utils.RespWithParamError(ctx, "id 只能为数字")
		return
	}
	err, resp := controller.CtrlMessageGet(uint(pid))
	utils.Resp(ctx, err, resp)
}
