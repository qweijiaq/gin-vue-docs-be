package jwts

import "github.com/dgrijalva/jwt-go/v4"

type JwtPayLoad struct {
	NickName string `json:"nickName"`
	RoleID   uint   `json:"roleID"`
	UserID   uint   `json:"userID"`
}

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}
