package user

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/plugins/log_stash"
	"gvd_server/service/common/response"
	"gvd_server/service/redis"
	"gvd_server/utils/jwts"
	"time"
)

// UserLogoutView 用户注销
// @Tags 用户管理
// @Summary 用户注销
// @Description 注销的登录
// @Param token header string true "token"
// @Router /api/logout [get]
// @Produce json
// @Success 200 {object} response.Response{}
func (UserApi) UserLogoutView(c *gin.Context) {
	log := log_stash.NewAction(c)
	token := c.Request.Header.Get("token")
	claims, _ := jwts.ParseToken(token)
	// 过期时间
	exp := claims.ExpiresAt
	// 距离过期时间还有多久
	diff := exp.Time.Sub(time.Now())
	// 设置一个具有过期时间的 key，它的过期时间就是 token 的过期时间
	err := redis.Logout(token, diff)
	if err != nil {
		global.Log.Error(err)
	}
	log.Info("用户注销成功")
	log.SetItemInfo("userID", claims.UserID)
	response.OKWithMsg("用户注销成功", c)
}
