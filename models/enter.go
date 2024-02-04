package models

import "time"

type Model struct {
	ID        uint      `gorm:"primaryKey" json:"id"`              // 主键 ID
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"` // 添加时间
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"` // 更新时间
}

type Pagination struct {
	Page  int    `json:"page" form:"page"`   // 页
	Limit int    `json:"limit" form:"limit"` // 限制 -- 即 1 页多少条
	Key   string `json:"key" form:"key"`     // 关键词 -- 搜索栏中输入的内容，通常用于模糊查询
	Sort  string `json:"sort" form:"sort"`   // 排序方式
}

// IDListRequest 批量操作 ID 列表对应的多个对象
type IDListRequest struct {
	IDList []uint `json:"idList" form:"idList" binding:"required" label:"ID列表"`
}

// IDRequest 单独操作 ID 对应的对象
type IDRequest struct {
	ID uint `json:"id" form:"id" uri:"id"`
}
