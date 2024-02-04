package api

import (
	"gvd_server/api/image"
	"gvd_server/api/logs"
	"gvd_server/api/user"
)

type Api struct {
	UserApi  user.UserApi
	ImageApi image.ImageApi
	LogApi   logs.LogApi
}

// App 指针化 Api 结构体，保证全局唯一性
var App = new(Api)
