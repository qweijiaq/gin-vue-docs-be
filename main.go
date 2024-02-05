package main

import (
	"gvd_server/core"
	_ "gvd_server/docs" // swag init 生成后的 docs 路径
	"gvd_server/flags"
	"gvd_server/global"
	"gvd_server/routers"
)

// @title 知识库项目 API 文档
// @version 1.0
// @description API 文档
// @host 127.0.0.1:3000
// @BasePath /
func main() {
	global.Log = core.InitLogger()
	global.Config = core.InitConfig()
	global.DB = core.InitMysql()
	global.Redis = core.InitRedis(0)

	option := flags.Parse()
	if option.Run() {
		return
	}

	router := routers.Routers()
	addr := global.Config.System.Addr()
	router.Run(addr)
}
