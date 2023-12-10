package user

import "github.com/gin-gonic/gin"

type UserCreateRequest struct {
	UserName string `json:"userName" binding:"required" label:"用户名"` // 用户名
	Password string `json:"password" binding:"required"`             // 密码
	NickName string `json:"nickName"`                                // 昵称
	RoleID   uint   `json:"roleID" binding:"required"`               // 角色 ID
}

func (UserApi) UserCreateView(c *gin.Context) {
	var cr UserCreateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		c.JSON(200, gin.H{"msg": "失败"})
		return
	}
	c.JSON(200, gin.H{"msg": "成功"})
	return
}
