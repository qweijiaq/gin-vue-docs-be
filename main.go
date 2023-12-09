package main

import (
	"fmt"
	"gvd_server/core"
	"gvd_server/global"
	"gvd_server/routers"
)

func main() {
	global.Log = core.InitLogger()
	global.Config = core.InitConfig()
	global.DB = core.InitMysql()
	global.Redis = core.InitRedis(0)

	val, err := global.Redis.Get("name").Result()
	fmt.Println(val, err)

	router := routers.Routers()
	addr := global.Config.System.Addr()
	router.Run(addr)
}
