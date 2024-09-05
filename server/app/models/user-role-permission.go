package models

import "gorm.io/gorm"

type UserRolePermission struct {
	ID        int            `json:"id" gorm:"primaryKey"` // 主键ID
	RoleID    int            `json:"roleId"`               // 绑定的角色ID
	Resource  string         `json:"resource"`             // 允许的资源
	Action    string         `json:"action"`               // 允许的行为
	DeletedAt gorm.DeletedAt `json:"-"`                    // 删除标记
}

type UserRolePermissionReq struct {
	ID       int    `json:"id" gorm:"primaryKey"` // 主键ID
	RoleID   int    `json:"roleId"`               // 绑定的角色ID
	Resource string `json:"resource"`             // 允许的资源
	Action   string `json:"action"`               // 允许的行为
}

type UserRolePermissionResp struct {
	ID        int            `json:"id" gorm:"primaryKey"` // 主键ID
	RoleID    int            `json:"roleId"`               // 绑定的角色ID
	Resource  string         `json:"resource"`             // 允许的资源
	Action    string         `json:"action"`               // 允许的行为
	DeletedAt gorm.DeletedAt `json:"-"`                    // 删除标记
}

type UserRolePermissionListResp struct {
	Total   int64                    `json:"total"`   // 总数
	Records []UserRolePermissionResp `json:"records"` // 数据列表
}
