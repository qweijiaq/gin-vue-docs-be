package models

/*
UserPwdDocModel 对于一个文档，如果通过该表可以查询到当前用户的 ID，则该用户不需要输入密码即可阅读该文档，否则需要输入密码
*/
type UserPwdDocModel struct {
	Model
	UserID uint `gorm:"column:user_id;comment:用户ID" json:"userID"` // 用户ID
	DocID  uint `gorm:"column:doc_id;comment:文档ID" json:"docID"`   // 文档ID
}
