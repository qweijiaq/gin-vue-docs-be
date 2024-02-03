package models

// LoginModel 用户登录数据
type LoginModel struct {
	Model
	UserID    uint      `gorm:"column:userID;comment:用户ID" json:"userID"` // 用户ID
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	IP        string    `gorm:"size:20;column:IP;comment:IP" json:"ip"`               // 登录的 IP
	Nickname  string    `gorm:"column:nickname;size:42;comment:用户昵称" json:"nickname"` // 用户昵称
	UA        string    `gorm:"size:256;column:UA;comment:UA" json:"ua"`              // User-Agent
	Token     string    `gorm:"size:256;column:token;comment:token" json:"token"`     // 登录时的 token
	Device    string    `gorm:"size:256;column:device;comment:device" json:"device"`  // 登录设备
	Addr      string    `gorm:"size:64;column:addr;comment:addr" json:"addr"`         // 登录地址
}
