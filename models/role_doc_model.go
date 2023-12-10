package models

type RoleDocModel struct {
	Model
	RoleID      uint      `gorm:"column:role_id;comment:角色id" json:"roleID"`
	RoleModel   RoleModel `gorm:"foreignKey:RoleID" json:"-"`
	DocID       uint      `gorm:"column:doc_id;comment:文档id" json:"docID"`
	DocModel    DocModel  `gorm:"foreignKey:DocID" json:"-"`
	Pwd         *string   `gorm:"column:pwd;comment:密码配置" json:"pwd"`                 // null ""  "有值"  优先级： 角色文档密码 > 角色密码
	FreeContent *string   `gorm:"column:freeContent;comment:试看配置" json:"freeContent"` // 试看部分 优先级：角色文档试看  > 文档试看字段 > 文档按照特殊字符分隔的试看
	Sort        int       `gorm:"column:sort;comment:排序" json:"sort"`                 // 排序
}
