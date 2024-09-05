package models

import (
	"edu-profit/utils"
	"gorm.io/gorm"
	"time"
)

type UserRole struct {
	ID        int            `json:"id" gorm:"primaryKey"`            // 主键ID
	RoleName  string         `json:"roleName"`                        // 角色名
	CreatedAt time.Time      `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time      `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
	DeletedAt gorm.DeletedAt `json:"-"`                               // 删除标记
}

type UserRoleReq struct {
	ID         int              `json:"id"`        // 主键ID
	RoleName   string           `json:"roleName"`  // 角色名
	DateRange  []time.Time      `json:"dateRange"` // 时间范围
	Sorted     string           `json:"sorted"`    // 排序
	OrderBy    string           `json:"orderBy"`   // 排序字段
	Pagination utils.Pagination `gorm:"embedded"`  // 分页
}

type UserRoleResp struct {
	ID         int                      `json:"id"`                                                // 主键ID
	RoleName   string                   `json:"roleName"`                                          // 角色名
	CreatedAt  time.Time                `json:"createdAt"`                                         // 创建时间
	UpdatedAt  time.Time                `json:"updatedAt"`                                         // 更新时间
	DeletedAt  gorm.DeletedAt           `json:"-"`                                                 // 删除标记
	Permission []UserRolePermissionResp `json:"permission" gorm:"foreignKey:RoleID;references:ID"` // 权限信息
}

type UserRoleListResp struct {
	Total   int64          `json:"total"`   // 总数
	Records []UserRoleResp `json:"records"` // 数据列表
}

func (UserRoleResp) TableName() string { return "user_role" }
