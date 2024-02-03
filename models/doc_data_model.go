package models

// DocDataModel 文档数据表
type DocDataModel struct {
	Model
	DocID     uint   `gorm:"column:docID;comment:文档ID" json:"docID"`        // 文档ID
	DocTitle  string `gorm:"column:docTitle;comment:文档标题" json:"docTitle"`  // 文档标题
	LookCount int    `gorm:"column:lookCount;comment:浏览量" json:"lookCount"` // 浏览量
	DiggCount int    `gorm:"column:diggCount;comment:点赞量" json:"diggCount"` // 点赞量
	CollCount int    `gorm:"column:collCount;comment:收藏量" json:"collCount"` // 收藏量
}
