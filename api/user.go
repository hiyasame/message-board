package api

import (
	"github.com/gin-gonic/gin"
	"message-board/controller"
	"message-board/model"
	"message-board/utils"
	"strconv"
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

func HandleSendVerifyCode(ctx *gin.Context) {
	email := ctx.PostForm("email")
	if utils.MatchEmailFormat(email) {
		utils.RespWithParamError(ctx, "邮箱地址不正确")
		return
	}
	err, resp := controller.CtrlSendVerifyCode(email)
	utils.Resp(ctx, err, resp)
}

func HandleChangePassword(ctx *gin.Context) {
	email := ctx.PostForm("email")
	password := ctx.PostForm("password")
	verify := ctx.PostForm("verify")
	if utils.MatchEmailFormat(email) {
		utils.RespWithParamError(ctx, "邮箱地址不正确")
		return
	}
	if utils.CheckPasswordStrength(password) {
		utils.RespWithParamError(ctx, "密码格式不正确")
		return
	}
	err, resp := controller.CtrlChangePassword(email, password, verify)
	utils.Resp(ctx, err, resp)
}

func HandleChangeUserDetail(ctx *gin.Context) {
	id, exist := ctx.Get("uid")
	if !exist {
		utils.RespWithParamError(ctx, "未登录")
		return
	}
	user := model.User{
		ID: uint(id.(int64)),
	}
	avatar, exist := ctx.GetPostForm("avatar")
	if exist {
		user.Avatar = avatar
	}
	bio, exist := ctx.GetPostForm("bio")
	if exist {
		user.Bio = bio
	}
	err, resp := controller.CtrlChangeDetail(&user)
	utils.Resp(ctx, err, resp)
}

func HandleGetUserDetail(ctx *gin.Context) {
	id, exist := ctx.GetPostForm("uid")
	if exist {
		uid, err := strconv.Atoi(id)
		if err != nil {
			utils.RespWithParamError(ctx, "uid必须为数字")
			return
		}
		uuid := uint(uid)
		err, resp := controller.CtrlGetDetail(&uuid, nil, nil)
		utils.Resp(ctx, err, resp)
		return
	}
	email, exist := ctx.GetPostForm("email")
	if exist {
		err, resp := controller.CtrlGetDetail(nil, nil, &email)
		utils.Resp(ctx, err, resp)
		return
	}
	name, exist := ctx.GetPostForm("name")
	if exist {
		err, resp := controller.CtrlGetDetail(nil, nil, &name)
		utils.Resp(ctx, err, resp)
		return
	}
}
