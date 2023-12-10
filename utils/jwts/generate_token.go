package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
	"gvd_server/global"
	"time"
)

// GenToken 生成 token
func GenToken(user JwtPayLoad) (string, error) {
	claims := CustomClaims{
		JwtPayLoad: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Duration(global.Config.Jwt.Expires) * time.Hour * 24)), // 过期时间
			Issuer:    global.Config.Jwt.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.Config.Jwt.Secret))
}
