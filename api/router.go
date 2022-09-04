package api

import (
	"github.com/gin-gonic/gin"
	"message-board/api/middleware"
	"net/http"
)

type HandlerFunctions = []gin.HandlerFunc

// Route 表示每一个路由
type Route struct {
	Name             string
	Method           string // it is the string for the HTTP method. ex) GET, POST etc..
	Pattern          string
	HandlerFunctions HandlerFunctions `json:"-"`
}

var Routes = []Route{
	{
		Name:             "注册",
		Method:           http.MethodPost,
		Pattern:          "/user/register",
		HandlerFunctions: HandlerFunctions{HandleUserRegister},
	},
	{
		Name:             "登录",
		Method:           http.MethodPost,
		Pattern:          "/user/login",
		HandlerFunctions: HandlerFunctions{HandleUserLogin},
	},
	{
		Name:             "发送验证码",
		Method:           http.MethodPost,
		Pattern:          "/user/verify",
		HandlerFunctions: HandlerFunctions{HandleSendVerifyCode},
	},
	{
		Name:             "修改密码",
		Method:           http.MethodPost,
		Pattern:          "/user/changepass",
		HandlerFunctions: HandlerFunctions{HandleChangePassword},
	},
	{
		Name:             "修改用户信息",
		Method:           http.MethodPost,
		Pattern:          "/user/detail",
		HandlerFunctions: HandlerFunctions{middleware.Auth(), HandleChangeUserDetail},
	},
	{
		Name:             "获取用户信息",
		Method:           http.MethodGet,
		Pattern:          "/user/detail",
		HandlerFunctions: HandlerFunctions{HandleGetUserDetail},
	},
	{
		Name:             "留言",
		Method:           http.MethodPut,
		Pattern:          "/message",
		HandlerFunctions: HandlerFunctions{middleware.Auth(), HandlePutMessage},
	},
	{
		Name:             "获取留言详情",
		Method:           http.MethodGet,
		Pattern:          "/message",
		HandlerFunctions: HandlerFunctions{HandleGetMessage},
	},
	{
		Name:             "回复指定条留言的回复",
		Method:           http.MethodPut,
		Pattern:          "/message/:id",
		HandlerFunctions: HandlerFunctions{middleware.Auth(), HandlePutChildMessage},
	},
	{
		Name:             "获取指定条留言的回复",
		Method:           http.MethodGet,
		Pattern:          "/message/:id",
		HandlerFunctions: HandlerFunctions{HandleGetChildrenMessage},
	},
}

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.Cors())
	InitRoutes(r)
	// By default, it serves on :8080 unless a
	// PORT environment variable was defined.
	err := r.Run()
	if err != nil {
		return
	}
}

func InitRoutes(engine *gin.Engine) {
	for _, v := range Routes {
		engine.Handle(v.Method, v.Pattern, v.HandlerFunctions...)
	}
}
