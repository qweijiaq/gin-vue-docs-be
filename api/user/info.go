package user

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/plugins/log_stash"
	"gvd_server/service/common/response"
	"gvd_server/utils/jwts"
)

type UserInfoResponse struct {
	models.UserModel
	UserName string `json:"userName"`
	Role     string `json:"role"`
}

// UserInfoView 用户信息
// @Tags 用户管理
// @Summary 用户信息
// @Description 获取用户的一些详细信息
// @Param token header string true "token"
// @Router /api/user_info [get]
// @Produce json
// @Success 200 {object} response.Response{data=UserInfoResponse}
func (UserApi) UserInfoView(c *gin.Context) {
	log := log_stash.NewAction(c)

	_claims, _ := c.Get("claims")
	claims, _ := _claims.(*jwts.CustomClaims)

	var user models.UserModel
	err := global.DB.Preload("RoleModel").Take(&user, claims.UserID).Error
	if err != nil {
		log.Error("获取用户信息失败")
		log.SetItemInfo("userID", claims.UserID)
		log.SetItemErr("失败原因", "该用户不存在")
		response.FailWithMsg("用户不存在", c)
		return
	}
	info := UserInfoResponse{
		UserModel: user,
		UserName:  user.UserName,
		Role:      user.RoleModel.Title,
	}
	response.OKWithData(info, c)
}
