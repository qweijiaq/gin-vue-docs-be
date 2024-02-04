package user

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/plugins/log_stash"
	"gvd_server/service/common/response"
	"gvd_server/utils/encryption"
)

// UserUpdateRequest 更新用户请求参数
type UserUpdateRequest struct {
	ID       uint   `json:"id" binding:"required" label:"用户ID"` // 用户ID
	Password string `json:"password" label:"密码"`                // 密码
	Nickname string `json:"nickname" label:"昵称"`                // 昵称
	RoleID   uint   `json:"roleID" label:"角色"`                  // 角色ID
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
	log := log_stash.NewAction(c)
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, c)
		return
	}

	var user models.UserModel
	err = global.DB.Take(&user, cr.ID).Error
	if err != nil {
		log.Error("用户更新失败")
		log.SetItemInfo("username", user.Username)
		log.SetItemErr("失败原因", "该用户不存在")
		response.FailWithMsg("该用户不存在", c)
		return
	}

	if cr.Password != "" {
		cr.Password = encryption.HashPwd(cr.Password)
	}
	if cr.RoleID != 0 {
		var role models.RoleModel
		err = global.DB.Take(&role, cr.RoleID).Error
		if err != nil {
			log.Error("用户更新失败")
			log.SetItemInfo("RoleID", cr.RoleID)
			log.SetItemErr("失败原因", "该角色不存在")
			response.FailWithMsg("该角色不存在", c)
			return
		}
	}

	err = global.DB.Model(&user).Updates(models.UserModel{
		Password: cr.Password,
		Nickname: cr.Nickname,
		RoleID:   cr.RoleID,
	}).Error
	if err != nil {
		global.Log.Error(err)
		log.Error("用户更新失败")
		log.SetItemInfo("username", user.Username)
		log.SetItemErr("失败原因", err.Error())
		response.FailWithMsg("用户更新失败", c)
		return
	}

	log.Info("用户更新成功")
	log.SetItemInfo("username", user.Username)
	response.OKWithMsg("用户更新成功", c)
}
