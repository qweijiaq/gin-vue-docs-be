package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) UserRouter() {
	app := api.App.UserApi

	router.POST("users", middleware.Auth(), app.UserCreateView) // 创建用户
	router.POST("login", app.UserLoginView)                     // 登录
}
