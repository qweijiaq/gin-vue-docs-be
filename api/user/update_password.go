package user

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/service/common/response"
	"gvd_server/utils/encryption"
	"gvd_server/utils/jwts"
)

type UserUpdatePasswordRequest struct {
	OldPwd   string `json:"oldPwd" binding:"required" label:"之前的密码"`
	Password string `json:"password" binding:"required" label:"新密码"`
}

// UserUpdatePasswordView 用户修改密码
// @Tags 用户管理
// @Summary 用户修改密码
// @Description 用户修改密码
// @Param token header string true "token"
// @Param data body UserUpdatePasswordRequest true "参数"
// @Router /api/user_password [put]
// @Produce json
// @Success 200 {object} response.Response{}
func (UserApi) UserUpdatePasswordView(c *gin.Context) {
	var cr UserUpdatePasswordRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithInValidError(err, &cr, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims, _ := _claims.(*jwts.CustomClaims)
	user, err := claims.GetUser()
	if err != nil {
		response.FailWithMsg("用户不存在", c)
		return
	}
	if !encryption.CheckPwd(user.Password, cr.OldPwd) {
		response.FailWithMsg("原密码错误", c)
		return
	}
	hashPwd := encryption.HashPwd(cr.Password)
	global.DB.Model(user).Update("password", hashPwd)

	response.OKWithMsg("用户密码修改成功", c)
	return

}
