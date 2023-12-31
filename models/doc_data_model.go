package models

// DocDataModel 文档数据表
type DocDataModel struct {
	Model
	DocID     uint   `gorm:"column:docID" json:"docID"`
	DocTitle  string `gorm:"column:docTitle" json:"docTitle"`
	LookCount int    `gorm:"column:lookCount" json:"lookCount"`
	DiggCount int    `gorm:"column:diggCount" json:"diggCount"`
	CollCount int    `gorm:"column:collCount" json:"collCount"`
}
