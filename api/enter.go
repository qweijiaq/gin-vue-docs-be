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

var App = new(Api)
