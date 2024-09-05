package models

import (
	"edu-profit/app/types"
	"gorm.io/gorm"
	"time"
)

type CourseClassification struct {
	ID        int              `json:"id" gorm:"primaryKey"`            // 主键ID
	Title     string           `json:"title"`                           // 分类标题
	Profit    int              `json:"profit"`                          // 每课时课时费
	Status    types.StatusType `json:"status" gorm:"default 1"`         // 状态 1:启用 2:冻结 3:删除
	CreatedAt time.Time        `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time        `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
	DeletedAt gorm.DeletedAt   `json:"-"`                               // 删除标记
}

type CourseClassificationReq struct {
	ID        int              `json:"id"`        // 主键ID
	Title     string           `json:"title"`     // 分类标题
	Profit    int              `json:"profit"`    // 每课时课时费
	Status    types.StatusType `json:"status"`    // 状态 1:启用 2:冻结 3:删除
	DateRange []time.Time      `json:"dateRange"` // 时间范围
}

type CourseClassificationResp struct {
	ID        int              `json:"id"`        // 主键ID
	Title     string           `json:"title"`     // 分类标题
	Profit    int              `json:"profit"`    // 每课时课时费
	Status    types.StatusType `json:"status"`    // 状态 1:启用 2:冻结 3:删除
	CreatedAt time.Time        `json:"createdAt"` // 创建时间
	UpdatedAt time.Time        `json:"updatedAt"` // 更新时间
	DeletedAt gorm.DeletedAt   `json:"-"`         // 删除标记
}

type CourseClassificationListResp struct {
	Total   int64                      `json:"total"`   // 总数
	Records []CourseClassificationResp `json:"records"` // 数据列表
}
