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
