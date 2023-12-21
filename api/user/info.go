package user

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
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
// @Description 用户信息
// @Param token header string true "token"
// @Router /api/user_info [get]
// @Produce json
// @Success 200 {object} response.Response{data=UserInfoResponse}
func (UserApi) UserInfoView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims, _ := _claims.(*jwts.CustomClaims)

	var user models.UserModel
	err := global.DB.Preload("RoleModel").Take(&user, claims.UserID).Error
	if err != nil {
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
