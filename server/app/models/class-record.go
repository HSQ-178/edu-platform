package models

import (
	"edu-profit/app/types"
	"gorm.io/gorm"
	"time"
)

type ClassRecord struct {
	ID        int64            `json:"id" gorm:"primaryKey"`            // 主键ID
	CourseID  int64            `json:"courseId"`                        // 课程ID
	UserID    int64            `json:"userId"`                          // 授课用户ID
	Status    types.StatusType `json:"status"`                          // 状态 1:签到 2:未签到 3:删除
	CreatedAt time.Time        `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time        `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
	DeletedAt gorm.DeletedAt   `json:"-"`                               // 删除标记
}

type ClassRecordReq struct {
	ID        int64            `json:"id"`        // 主键ID
	CourseID  int64            `json:"courseId"`  // 课程ID
	UserID    int64            `json:"userId"`    // 授课用户ID
	Status    types.StatusType `json:"status"`    // 状态 1:签到 2:未签到 3:删除
	DateRange []time.Time      `json:"dateRange"` // 时间范围
}

type ClassRecordResp struct {
	ID        int64            `json:"id"`        // 主键ID
	CourseID  int64            `json:"courseId"`  // 课程ID
	UserID    int64            `json:"userId"`    // 授课用户ID
	Status    types.StatusType `json:"status"`    // 状态 1:签到 2:未签到 3:删除
	CreatedAt time.Time        `json:"createdAt"` // 创建时间
	UpdatedAt time.Time        `json:"updatedAt"` // 更新时间
	DeletedAt gorm.DeletedAt   `json:"-"`         // 删除标记
	Course    CourseResp       `json:"course"`    // 课程信息
	User      UserResp         `json:"user"`      // 授课用户信息
}

type ClassRecordListResp struct {
	Total   int64             `json:"total"`   // 总数
	Records []ClassRecordResp `json:"records"` // 数据列表
}
