package models

type UserCollectDocModel struct {
	Model
	DocID     uint      `gorm:"column:doc_id" json:"docID"`
	DocModel  DocModel  `gorm:"foreignKey:DocID"`
	UserID    uint      `gorm:"column:user_id" json:"userID"`
	UserModel UserModel `gorm:"foreignKey:UserID"`
}
