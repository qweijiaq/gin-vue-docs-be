package list

import (
	"fmt"
	"gorm.io/gorm"
	"gvd_server/global"
	"gvd_server/models"
)

type Option struct {
	models.Pagination          // 分页查询
	Likes             []string // 需要模糊匹配的字段列表
	Debug             bool     // 是否打印 SQL
	Where             *gorm.DB // 条件查询
	Preload           []string // 预加载的字段列表
}

// QueryList 分页列表查询
func QueryList[T any](model T, option Option) (list []T, count int, err error) {
	// 查 model 中非空字段
	query := global.DB.Where(model)
	if option.Debug {
		query = query.Debug()
	}

	// 默认按时间倒序排列
	if option.Sort == "" {
		option.Sort = "createdAt desc"
	}

	// 默认一页显示 10 条
	if option.Limit <= 0 {
		option.Limit = 10
	}
	// 如果有高级查询就加上
	if option.Where != nil {
		query.Where(option.Where)
	}

	// 模糊匹配
	if option.Key != "" {
		likeQuery := global.DB.Where("")
		for index, column := range option.Likes {
			// 第一个模糊匹配和前面的关系是 and 关系，后面的和前面的模糊匹配是 or 的关系
			if index == 0 {
				// 这里注意要防止 SQL 注入 -- 最简单的方式就是用 ? 而不是直接字符串拼接
				likeQuery.Where(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			} else {
				likeQuery.Or(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			}
		}
		// 整个模糊匹配是一个整体
		query = query.Where(likeQuery)
	}

	// 查列表，获取总数
	count = int(query.Find(&list).RowsAffected)

	// 预加载
	for _, preload := range option.Preload {
		query = query.Preload(preload)
	}

	// 计算偏移
	offset := (option.Page - 1) * option.Limit

	err = query.Limit(option.Limit).
		Offset(offset).
		Order(option.Sort).Find(&list).Error

	return
}
