/*
	对于一个文档，如果通过该表可以查询到当前用户的 ID，则该用户不需要输入密码即可阅读该文档，否则需要输入密码
*/

package models

type UserPwdDocModel struct {
	Model
	UserID uint `gorm:"column:user_id" json:"userID"`
	DocID  uint `gorm:"column:doc_id" json:"docID"`
}
