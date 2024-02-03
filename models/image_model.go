package models

import "fmt"

type ImageModel struct {
	Model
	UserID    uint      `gorm:"column:userID;comment:用户ID" json:"userID"` // 用户ID
	UserModel UserModel `gorm:"foreignKey:UserID" json:"-"`
	FileName  string    `gorm:"column:fileName;size:64；comment:文件名" json:"fileName"` // 文件名
	Size      int64     `gorm:"column:size;comment:文件大小，单位字节" json:"size"`           // 文件大小，单位字节
	Path      string    `gorm:"column:path;size:128;comment:文件路径" json:"path"`       // 文件路径
	Hash      string    `gorm:"column:hash;size:64;comment:文件的hash" json:"hash"`     // 文件的hash -- 避免上传重复图片
}

// WebPath 便于线上访问图片时拼接路径
func (image ImageModel) WebPath() string {
	return fmt.Sprintf("/%s", image.Path)
}
