package user

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/response"
	"gvd_server/utils/encryption"
	"gvd_server/utils/jwts"
	"time"
)

type UserLoginRequest struct {
	UserName string `json:"userName" binding:"required" label:"用户名"`
	Password string `json:"password" binding:"required" label:"密码"`
}

func (UserApi) UserLoginView(c *gin.Context) {
	var cr UserLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithInValidError(err, &cr, c)
		return
	}

	var user models.UserModel
	err = global.DB.Take(&user, "userName = ?", cr.UserName).Error
	if err != nil {
		global.Log.Warn("用户名不存在", cr.UserName)
		response.FailWithMsg("用户名或密码错误", c)
		return
	}
	if !encryption.CheckPwd(user.Password, cr.Password) {
		global.Log.Warn("用户密码错误", cr.UserName, cr.Password)
		response.FailWithMsg("用户名或密码错误", c)
		return
	}

	token, err := jwts.GenToken(jwts.JwtPayLoad{
		NickName: user.NickName,
		RoleID:   user.RoleID,
		UserID:   user.ID,
	})
	if err != nil {
		global.Log.Error(err)
		response.FailWithMsg("生成token失败", c)
		return
	}
	c.Request.Header.Set("token", token)
	global.DB.Model(&user).Update("lastLogin", time.Now())

	_ip := c.ClientIP()
	addr := c.RemoteIP()
	ua := c.Request.Header.Get("User-Agent")

	go func() {
		// 加一条登录记录
		err = global.DB.Create(&models.LoginModel{
			UserID:   user.ID,
			IP:       _ip,
			NickName: user.NickName,
			UA:       ua,
			Token:    token,
			Addr:     addr,
		}).Error

		if err != nil {
			global.Log.Error(err)
		}
	}()
	response.OKWithData(token, c)
	return

}