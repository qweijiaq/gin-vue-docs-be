package log_stash

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type JwtPayLoad struct {
	NickName string `json:"nickName"`
	RoleID   uint   `json:"roleID"`
	UserID   uint   `json:"userID"`
	UserName string `json:"userName"`
}

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

func parseToken(token string) (jwtPayload *JwtPayLoad) {
	Token, _ := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(""), nil
	})
	if Token == nil || Token.Claims == nil {
		return nil
	}
	claims, ok := Token.Claims.(*CustomClaims)
	if !ok {
		return nil
	}
	return &claims.JwtPayLoad
}
