package models

type DocModel struct {
	Model
	Title           string      `gorm:"comment:文档标题;size:32" json:"title"`                                                     // 标题
	Content         string      `gorm:"comment:文档内容" json:"-"`                                                                 // 内容
	DiggCount       int         `gorm:"comment:点赞量;column:diggCount" json:"diggCount"`                                         // 点赞量
	LookCount       int         `gorm:"comment:浏览量;column:lookCount" json:"lookCount"`                                         // 收藏量
	Key             string      `gorm:"comment:key;not null;unique" json:"key"`                                                // 用于寻找文档上下级关系
	ParentID        *uint       `gorm:"comment:父文档ID;column:parentID" json:"parentID"`                                         // 父文档ID
	ParentModel     *DocModel   `gorm:"foreignKey:ParentID" json:"-"`                                                          // 父文档
	Child           []*DocModel `gorm:"foreignKey:ParentID" json:"child"`                                                      // 子孙文档
	FreeContent     string      `gorm:"comment:预览部分;column:freeContent" json:"freeContent"`                                    // 预览部分
	UserCollDocList []UserModel `gorm:"many2many:user_collect_doc_models;joinForeignKey:DocID;JoinReferences:UserID" json:"-"` // 收藏该文档的用户列表
}
