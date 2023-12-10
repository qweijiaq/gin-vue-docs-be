package models

import "time"

type UserModel struct {
	Model
	UserName  string    `gorm:"column:userName;size:36;unique;not null;comment:用户名" json:"userName"` // 用户名
	Password  string    `gorm:"column:password;size:128;comment:密码"  json:"-"`                       // 密码
	Avatar    string    `gorm:"column:avatar;size:256;comment:头像"  json:"avatar"`                    // 头像
	NickName  string    `gorm:"column:nickName;size:36;comment:昵称"  json:"nickName"`                 // 昵称
	Email     string    `gorm:"column:email;size:128;comment:邮箱"  json:"email"`                      // 邮箱
	Token     string    `gorm:"column:token;size:64;comment:其他平台的唯一id"  json:"-"`                    // 其他平台的唯一 ID
	IP        string    `gorm:"column:ip;size:16;comment:ip地址"  json:"ip"`                           // IP
	Addr      string    `gorm:"column:addr;size:64;comment:地址"  json:"addr"`                         // 地址
	RoleID    uint      `gorm:"column:roleID;comment:用户对应的角色" json:"roleID"`                         // 用户对应的角色
	LastLogin time.Time `gorm:"column:lastLogin" json:"lastLogin"`                                   // 用户最后登录时间
	RoleModel RoleModel `gorm:"foreignKey:RoleID" json:"roleModel"`                                  // 用户角色信息
}
