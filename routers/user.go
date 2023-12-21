package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) UserRouter() {
	app := api.App.UserApi

	router.POST("users", middleware.JwtAdmin(), app.UserCreateView)               // 创建用户
	router.POST("login", app.UserLoginView)                                       // 登录
	router.PUT("users", middleware.JwtAdmin(), app.UserUpdateView)                // 管理员更新用户
	router.GET("users", middleware.JwtAdmin(), app.UserListView)                  // 用户列表
	router.DELETE("users", middleware.JwtAdmin(), app.UserRemoveView)             // 删除用户
	router.GET("logout", middleware.JwtAuth(), app.UserLogoutView)                // 用户注销
	router.GET("user_info", middleware.JwtAuth(), app.UserInfoView)               // 用户详情信息
	router.PUT("user_password", middleware.JwtAuth(), app.UserUpdatePasswordView) // 用户修改密码
	router.PUT("user_info", middleware.JwtAuth(), app.UserUpdateInfoView)         // 用户修改信息
}
