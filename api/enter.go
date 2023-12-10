package api

import "gvd_server/api/user"

type Api struct {
	UserApi user.UserApi
}

var App = new(Api)
