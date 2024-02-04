package encryption

import (
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// HashPwd 将密码 hash
func HashPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		logrus.Errorln(err)
	}
	return string(hash)
}

// CheckPwd 验证密码   第一个参数是 hash 之后的密码，第二个参数是u需要被验证的密码
func CheckPwd(hashedPwd string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
	if err != nil {
		return false
	}
	return true
}
