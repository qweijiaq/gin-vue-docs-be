package models

// UserCollectDocModel 这个表是一个自定义的连接表，用于关联用户表和文档表
type UserCollectDocModel struct {
	Model
	DocID     uint      `gorm:"comment:文档ID;column:docID" json:"docID"` // 文档ID
	DocModel  DocModel  `gorm:"foreignKey:DocID"`
	UserID    uint      `gorm:"comment:用户ID;column:userID" json:"userID"` // 用户ID
	UserModel UserModel `gorm:"foreignKey:UserID"`
}
