package services

import (
	"edu-profit/app/types"
	"edu-profit/utils"
	"gorm.io/gorm"
	"time"
)

type QueryOption func(*gorm.DB)

func ApplyFilters(db *gorm.DB, opts ...QueryOption) {
	for _, opt := range opts {
		opt(db) // 调用每个过滤器，应用到db上
	}
}

func WithID(id int64) QueryOption {
	return func(db *gorm.DB) {
		if id != 0 {
			db.Where("id = ?", id)
		}
	}
}

func WithRoleID(roleID int) QueryOption {
	return func(db *gorm.DB) {
		if roleID != 0 {
			db.Where("role_id = ?", roleID)
		}
	}
}

func WithUsername(username string) QueryOption {
	return func(db *gorm.DB) {
		if username != "" {
			db.Where("username Like ?", "%"+username+"%")
		}
	}
}

func WithNickname(nickname string) QueryOption {
	return func(db *gorm.DB) {
		if nickname != "" {
			db.Where("nickname Like ?", "%"+nickname+"%")
		}
	}
}

func WithEmail(email string) QueryOption {
	return func(db *gorm.DB) {
		if email != "" {
			db.Where("email = ?", email)
		}
	}
}

func WithPhone(phone string) QueryOption {
	return func(db *gorm.DB) {
		if phone != "" {
			db.Where("phone = ?", phone)
		}
	}
}

func WithStatus(status types.StatusType) QueryOption {
	return func(db *gorm.DB) {
		if status != 0 {
			db.Where("status = ?", status)
		}
	}
}

func WithDateRange(dateRange []time.Time) QueryOption {
	return func(db *gorm.DB) {
		if len(dateRange) <= 2 {
			if !dateRange[0].IsZero() {
				db.Where("created_at > ?", dateRange[0])
			}
			if !dateRange[1].IsZero() {
				db.Where("created_at < ?", dateRange[1])
			}
		}
	}
}

func WithPagination(pagination utils.Pagination) QueryOption {
	return func(db *gorm.DB) {
		if pagination.Page > 0 && pagination.PageSize > 0 {
			db.Scopes(utils.Paginate(&pagination))
		}
	}
}
