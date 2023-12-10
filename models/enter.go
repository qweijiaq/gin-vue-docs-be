package models

import "time"

type Model struct {
	ID        uint      `json:"id" gorm:"primaryKey"`              // 主键 ID
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"` // 添加时间
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"` // 更新时间
}
