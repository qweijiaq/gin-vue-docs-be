package user

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/response"
	"gvd_server/utils/encryption"
)

type UserUpdateRequest struct {
	ID       uint   `json:"id" binding:"required" label:"用户id"`
	Password string `json:"password"` // 密码
	NickName string `json:"nickName"` // 昵称
	RoleID   uint   `json:"roleID"`   // 角色id
}

// UserUpdateView 管理员更新用户信息
// @Tags 用户管理
// @Summary 管理员更新用户信息
// @Description 管理员更新用户的一些信息
// @Param token header string true "token"
// @Param data body UserUpdateRequest true "参数"
// @Router /api/users [put]
// @Produce json
// @Success 200 {object} response.Response{}
func (UserApi) UserUpdateView(c *gin.Context) {
	var cr UserUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	var user models.UserModel
	err = global.DB.Take(&user, cr.ID).Error
	if err != nil {
		response.FailWithMsg("用户不存在", c)
		return
	}

	if cr.Password != "" {
		cr.Password = encryption.HashPwd(cr.Password)
	}
	if cr.RoleID != 0 {
		var role models.RoleModel
		err = global.DB.Take(&role, cr.RoleID).Error
		if err != nil {
			response.FailWithMsg("角色不存在", c)
			return
		}
	}

	err = global.DB.Model(&user).Updates(models.UserModel{
		Password: cr.Password,
		NickName: cr.NickName,
		RoleID:   cr.RoleID,
	}).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMsg("用户更新失败", c)
		return
	}

	response.OKWithMsg("用户更新成功", c)

}
