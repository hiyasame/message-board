package api

import (
	"github.com/gin-gonic/gin"
	"message-board/controller"
	"message-board/utils"
)

func HandleUserRegister(ctx *gin.Context) {
	name := ctx.PostForm("name")
	email := ctx.PostForm("email")
	verify := ctx.PostForm("verify")
	password := ctx.PostForm("password")
	if len(name) > 32 {
		utils.RespWithParamError(ctx, "名称太长了")
		return
	}
	if utils.CheckName(name) {
		utils.RespWithParamError(ctx, "名称格式不支持")
		return
	}
	if utils.MatchEmailFormat(email) {
		utils.RespWithParamError(ctx, "邮箱地址不正确")
		return
	}
	if err := utils.VerifyInputCode(email, verify); err != nil {
		utils.RespWithParamError(ctx, "验证码不正确")
		return
	}
	if utils.CheckPasswordStrength(password) {
		utils.RespWithParamError(ctx, "密码格式不正确")
		return
	}
	err, resp := controller.CtrlUserRegister(name, email, password)
	utils.Resp(ctx, err, resp)
}

func HandleUserLogin(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	err, resp := controller.CtrlUserLogin(email, password)
	utils.Resp(ctx, err, resp)
}
