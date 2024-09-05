package models

import (
	"edu-profit/app/types"
	"time"
)

type Course struct {
	ID                     int64            `json:"id" gorm:"primaryKey"`            // 主键ID
	CourseClassificationID int              `json:"courseClassificationId"`          // 课程分类ID
	UserID                 int64            `json:"userId"`                          // 用户ID
	StudentName            string           `json:"studentName"`                     // 学生名
	Multiple               int              `json:"multiple" gorm:"default 2"`       // 是否多名学生 1:是 2:否
	Status                 types.StatusType `json:"status" gorm:"default 1"`         // 状态 1:启用 2:冻结 3:删除
	CreatedAt              time.Time        `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt              time.Time        `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

type CourseReq struct {
	IDStr                  string           `json:"id"`                     // 主键IDStr
	UserIDStr              string           `json:"userIdStr"`              // 用户IDStr
	ID                     int64            `json:"-"`                      // 主键ID
	CourseClassificationID int              `json:"courseClassificationId"` // 课程分类ID
	UserID                 int64            `json:"userId"`                 // 用户ID
	StudentName            string           `json:"studentName"`            // 学生名
	Multiple               int              `json:"multiple"`               // 是否多名学生 1:是 2:否
	Status                 types.StatusType `json:"status"`                 // 状态 1:启用 2:冻结 3:删除
	DateRange              []time.Time      `json:"dateRange"`              // 时间范围
}

type CourseResp struct {
	ID                     int64                    `json:"id"`                     // 主键ID
	CourseClassificationID int                      `json:"courseClassificationId"` // 课程分类ID
	UserID                 int64                    `json:"userId"`                 // 用户ID
	StudentName            string                   `json:"studentName"`            // 学生名
	Multiple               int                      `json:"multiple"`               // 是否多名学生 1:是 2:否
	Status                 types.StatusType         `json:"status"`                 // 状态 1:启用 2:冻结 3:删除
	CreatedAt              time.Time                `json:"createdAt"`              // 创建时间
	UpdatedAt              time.Time                `json:"updatedAt"`              // 更新时间
	CourseClassification   CourseClassificationResp `json:"courseClassification"`   // 课程分类信息
}

type CourseListResp struct {
	Total   int64        `json:"total"`   // 总数
	Records []CourseResp `json:"records"` // 数据列表
}
