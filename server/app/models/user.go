package models

import (
	"edu-profit/app/types"
	"edu-profit/utils"
	"time"
)

type User struct {
	ID        int64            `json:"id" gorm:"primaryKey"`            // 主键ID
	RoleID    int              `json:"roleId"`                          // 角色ID
	Username  string           `json:"username"`                        // 用户名
	Password  string           `json:"password"`                        // 密码
	Name      string           `json:"name"`                            // 姓名
	Nickname  string           `json:"nickname"`                        // 昵称
	Email     string           `json:"email" gorm:"default NULL"`       // 邮箱
	Phone     string           `json:"phone" gorm:"default NULL"`       // 手机号
	Avatar    string           `json:"avatar" gorm:"default NULL"`      // 头像
	Status    types.StatusType `json:"status" gorm:"default 1"`         // 状态 1:启用 2:冻结 3:删除
	CreatedAt time.Time        `json:"createdAt" gorm:"autoCreateTime"` // 创建时间
	UpdatedAt time.Time        `json:"updatedAt" gorm:"autoUpdateTime"` // 更新时间
}

type UserReq struct {
	IDStr      string           `json:"id"`        // 主键ID Str
	ID         int64            `json:"-"`         // 主键ID
	RoleID     int              `json:"roleId"`    // 角色ID
	Username   string           `json:"username"`  // 用户名
	Password   string           `json:"password"`  // 密码
	Name       string           `json:"name"`      // 姓名
	Nickname   string           `json:"nickname"`  // 昵称
	Email      string           `json:"email"`     // 邮箱
	Phone      string           `json:"phone"`     // 手机号
	Avatar     string           `json:"avatar"`    // 头像
	Status     types.StatusType `json:"status"`    // 状态 1:启用 2:冻结 3:删除
	DateRange  []time.Time      `json:"dateRange"` // 时间范围
	Sorted     string           `json:"sorted"`    // 排序
	OrderBy    string           `json:"orderBy"`   // 排序字段
	Pagination utils.Pagination `gorm:"embedded"`  // 分页
}

type UserResp struct {
	ID        int64            `json:"id"`                                              // 主键ID
	RoleID    int              `json:"roleId"`                                          // 角色ID
	Username  string           `json:"username"`                                        // 用户名
	Password  string           `json:"-"`                                               // 密码
	Name      string           `json:"name"`                                            // 姓名
	Nickname  string           `json:"nickname"`                                        // 昵称
	Email     string           `json:"email"`                                           // 邮箱
	Phone     string           `json:"phone"`                                           // 手机号
	Avatar    string           `json:"avatar"`                                          // 头像
	Status    types.StatusType `json:"status"`                                          // 状态 1:启用 2:冻结 3:删除
	CreatedAt time.Time        `json:"createdAt"`                                       // 创建时间
	UpdatedAt time.Time        `json:"updatedAt"`                                       // 更新时间
	UserRole  UserRoleResp     `json:"userRole" gorm:"foreignKey:ID;references:RoleID"` // 用户角色信息
}

type UserListResp struct {
	Total   int64      `json:"total"`   // 总数
	Records []UserResp `json:"records"` // 数据列表
}

type UserRegisterReq struct {
	Username string `json:"username"` // 用户名
	Password string `json:"password"` // 密码
	Nickname string `json:"nickname"` // 昵称
}

type UserLoginReq struct {
	Type     int    `json:"type"`     // 1: 用户名/手机号/邮箱 + 密码  2: 手机号 + 验证码  3: 邮箱 + 验证码
	Username string `json:"username"` // 用户名/手机号/邮箱
	Password string `json:"password"` // 密码
}

type UserLoginResp struct {
	Token string   `json:"token"` // token
	User  UserResp `json:"user"`  // 用户信息
}

func (UserResp) TableName() string {
	return "user"
}
