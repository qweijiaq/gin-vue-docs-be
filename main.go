package main

import (
	"gvd_server/core"
	"gvd_server/global"
	"gvd_server/routers"
)

func main() {
	global.Log = core.InitLogger()
	global.Config = core.InitConfig()
	global.DB = core.InitMysql()
	global.Redis = core.InitRedis(0)

	// option := flags.Parse()

	router := routers.Routers()
	addr := global.Config.System.Addr()
	router.Run(addr)
}
