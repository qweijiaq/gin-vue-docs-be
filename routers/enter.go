package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"gvd_server/global"
	"gvd_server/middleware"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func Routers() *gin.Engine {
	router := gin.Default()

	if global.Config.System.Env == "dev" {
		router.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
	}

	// 创建一个以 api 开头的路由分组
	apiGroup := router.Group("api")
	apiGroup.Use(middleware.LogMiddleWare())
	// 又将这个路由分组赋给了 RouterGroup
	routerGroup := RouterGroup{
		apiGroup,
	}

	// 第一个参数是 web 的访问别名  第二个参数是内部的映射目录
	// 线上如果有 Nginx，这一步可以省略
	router.Static("/uploads", "uploads")

	routerGroup.UserRouter()
	routerGroup.ImageRouter()

	return router
}
