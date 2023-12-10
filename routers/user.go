package routers

import "gvd_server/api"

func (router RouterGroup) UserRouter() {
	app := api.App.UserApi

	router.POST("users", app.UserCreateView)
}
