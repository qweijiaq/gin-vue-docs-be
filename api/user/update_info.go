package user

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/plugins/log_stash"
	"gvd_server/service/common/response"
	"gvd_server/utils/jwts"
)

// UserUpdateInfoRequest 用户更新自己的信息时传入的参数
type UserUpdateInfoRequest struct {
	Nickname string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 头像
}

// UserUpdateInfoView 用户更新自己的信息
// @Tags 用户管理
// @Summary 用户更新个人信息
// @Description 用户自己更新个人的一些信息
// @Param token header string true "token"
// @Param data body UserUpdateInfoRequest true "参数"
// @Router /api/user_info [put]
// @Produce json
// @Success 200 {object} response.Response{}
func (UserApi) UserUpdateInfoView(c *gin.Context) {
	var cr UserUpdateInfoRequest
	log := log_stash.NewAction(c)
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		response.FailWithError(err, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims, _ := _claims.(*jwts.CustomClaims)
	user, err := claims.GetUser()
	if err != nil {
		log.Error("用户信息修改失败")
		log.SetItemInfo("username", user.Username)
		log.SetItemErr("失败原因", "该用户不存在")
		response.FailWithMsg("该用户不存在", c)
		return
	}

	if cr.Avatar != "" {
		var image models.ImageModel
		err = global.DB.Take(&image, "userID = ? and path = ?", claims.UserID, cr.Avatar[1:]).Error
		if err != nil {
			log.Error("用户信息修改失败")
			log.SetItemErr("失败原因", "用户头像不存在")
			response.FailWithMsg("用户头像不存在", c)
			return
		}
	}
	if !(cr.Nickname == "" && cr.Avatar == "") {
		global.DB.Model(user).Updates(models.UserModel{
			Avatar:   cr.Avatar,
			Nickname: cr.Nickname,
		})
	}
	log.Info("用户信息修改成功")
	log.SetItemInfo("username", user.Username)
	response.OKWithMsg("用户信息修改成功", c)
	return
}
