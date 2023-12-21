package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/response"
	"gvd_server/utils/encryption"
	"time"
)

type UserCreateRequest struct {
	UserName string `json:"userName" binding:"required" label:"用户名"` // 用户名
	Password string `json:"password" binding:"required"`             // 密码
	NickName string `json:"nickName"`                                // 昵称
	RoleID   uint   `json:"roleID" binding:"required"`               // 角色 ID
}

// UserCreateView 创建用户
// @Tags 用户管理
// @Summary 创建用户
// @Description 创建用户，只能管理员创建
// @Param data body UserCreateRequest true "参数"
// @Param token header string true "token"
// @Router /api/users [post]
// @Produce json
// @Success 200 {object} response.Response{}
func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithInValidError(err, &cr, c)
		return
	}
	var user models.UserModel
	err = global.DB.Take(&user, "userName = ?", cr.UserName).Error
	if err == nil {
		response.FailWithMsg("该用户名已存在", c)
		return
	}
	if cr.NickName == "" {
		// 昵称不存在, 按规律生成
		var maxID uint
		err = global.DB.Model(models.UserModel{}).Select("max(id)").Scan(&maxID).Error
		if err != nil {
			maxID = 0
		}
		user.NickName = fmt.Sprintf("用户_%d", maxID+1)
	}
	var role models.RoleModel
	err = global.DB.Take(&role, cr.RoleID).Error
	if err != nil {
		response.FailWithMsg("该角色不存在", c)
		return
	}
	user.Password = encryption.HashPwd(user.Password)
	err = global.DB.Create(&models.UserModel{
		UserName:  cr.UserName,
		Password:  cr.Password,
		NickName:  cr.NickName,
		IP:        c.RemoteIP(),
		RoleID:    cr.RoleID,
		LastLogin: time.Now(),
	}).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMsg("用户创建失败", c)
		return
	}

	response.OKWithMsg("用户创建成功", c)
	return
}
