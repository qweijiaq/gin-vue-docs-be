package routers

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	*gin.RouterGroup
}

func Routers() *gin.Engine {
	router := gin.Default()

	// 创建一个以 api 开头的路由分组
	apiGroup := router.Group("api")
	// 又将这个路由分组赋给了 RouterGroup
	routerGroup := RouterGroup{
		apiGroup,
	}

	routerGroup.UserRouter()

	return router
}
