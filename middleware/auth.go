package middleware

import (
	"github.com/gin-gonic/gin"
	"gvd_server/service/common/response"
	"gvd_server/service/redis"
	"gvd_server/utils/jwts"
)

// JwtAuth 验证用户是否登录
func JwtAuth() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			response.FailWithMsg("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwts.ParseToken(token)
		if err != nil {
			response.FailWithMsg("token错误", c)
			c.Abort()
			return
		}

		ok := redis.CheckLogout(token)
		if ok {
			response.FailWithMsg("token已注销", c)
			c.Abort()
			return
		}

		c.Set("claims", claims)
	}
}
