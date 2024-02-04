package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
	"gvd_server/global"
	"gvd_server/models"
)

// JwtPayLoad 扩展的 JWT 负载
type JwtPayLoad struct {
	Nickname string `json:"nickname"`
	RoleID   uint   `json:"roleID"`
	UserID   uint   `json:"userID"`
}

// CustomClaims 自定义 Claims
type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims // 标准（默认）Claims
}

// GetUser 封装一个从 token 中查询用户的方法
func (c CustomClaims) GetUser() (user *models.UserModel, err error) {
	user = new(models.UserModel)
	err = global.DB.Take(user, c.UserID).Error
	return
}
