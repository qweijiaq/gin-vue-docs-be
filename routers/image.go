package routers

import (
	"gvd_server/api"
	"gvd_server/middleware"
)

func (router RouterGroup) ImageRouter() {
	app := api.App.ImageApi

	router.POST("images", middleware.JwtAuth(), app.ImageUploadView) // 上传图片
	router.GET("images", middleware.JwtAdmin(), app.ImageListView)   // 图片列表
}
