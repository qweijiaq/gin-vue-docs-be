package user

import (
	"github.com/gin-gonic/gin"
	"gvd_server/models"
	"gvd_server/service/common/list"
	"gvd_server/service/common/response"
)

// UserListRequest 获取用户列表时的请求参数
type UserListRequest struct {
	models.Pagination
	RoleID uint `json:"roleID" form:"roleID"` // 角色 ID -- 可根据角色查询对应的用户列表
}

// UserListView 用户列表
// @Tags 用户管理
// @Summary 用户列表
// @Description 获取用户列表
// @Param data query UserListRequest true "参数"
// @Param token header string true "token"
// @Router /api/users [get]
// @Produce json
// @Success 200 {object} response.Response{data=response.ListResponse[models.UserModel]}
func (UserApi) UserListView(c *gin.Context) {
	var cr UserListRequest
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		response.FailWithInValidError(err, &cr, c)
		return
	}
	_list, count, _ := list.QueryList(models.UserModel{RoleID: cr.RoleID}, list.Option{
		Pagination: cr.Pagination,                    // 分页
		Likes:      []string{"nickname", "username"}, // 根据昵称和用户名实现模糊查询
		Preload:    []string{"RoleModel"},            // 预加载出用户对应的角色
	})
	response.OKWithList(_list, count, c)
	return
}
