package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
	"gvd_server/global"
)

// ParseToken token 解析
func ParseToken(token string) (*CustomClaims, error) {
	Token, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(global.Config.Jwt.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := Token.Claims.(*CustomClaims)
	if !ok {
		// 数据不一致
		return nil, err
	}
	if !Token.Valid {
		// 令牌无效
		return nil, err
	}
	return claims, nil
}
