package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"gvd_server/global"
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
	// 又将这个路由分组赋给了 RouterGroup
	routerGroup := RouterGroup{
		apiGroup,
	}

	routerGroup.UserRouter()
	routerGroup.ImageRouter()

	return router
}
