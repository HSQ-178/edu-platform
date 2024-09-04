package utils

import "gorm.io/gorm"

type Pagination struct {
	Page     int `form:"page" json:"page"`
	PageSize int `form:"pageSize" json:"pageSize"`
	Offset   int `form:"offset" json:"offset"`
}

func Paginate(pagination *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := pagination.Page
		if page == 0 {
			page = 1
		}
		pageSize := pagination.PageSize
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := pagination.Offset

		if pagination.Offset == 0 {
			offset = (page - 1) * pageSize
		}

		return db.Offset(offset).Limit(pageSize)
	}
}
