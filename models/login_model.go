package models

// LoginModel 用户登录数据
type LoginModel struct {
	Model
	UserID    uint      `gorm:"column:userID" json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	IP        string    `gorm:"size:20" json:"ip"` // 登录的 IP
	NickName  string    `gorm:"column:nickName;size:42" json:"nickName"`
	UA        string    `gorm:"size:256" json:"ua"` // UA
	Token     string    `gorm:"size:256" json:"token"`
	Device    string    `gorm:"size:256" json:"device"` // 登录设备
	Addr      string    `gorm:"size:64" json:"addr"`
}
