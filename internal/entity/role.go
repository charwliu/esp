package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Roles []*Role

// Role 角色信息表
type Role struct {
	ID            int64          `gorm:"primary_key;not null" json:"roleId"` // 角色ID
	RoleName      string         `gorm:"size:32;not null" json:"roleName"`   // 角色名称
	RoleKey       string         `gorm:"size:100;not null" json:"roleKey"`   // 角色权限字符串
	RoleSort      int            `gorm:"not null" json:"roleSort"`           // 显示顺序
	DataScope     string         `gorm:"size:1;default:1" json:"dataScope"`  // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Status        string         `gorm:"size:1;not null" json:"status"`      // 角色状态（0正常 1停用）
	UpdateTime    time.Time      `json:"updateTime"`                         // 更新时间
	Remark        string         `gorm:"size:500" json:"remark"`             // 备注
	CreatedBy     string         `gorm:"size:64;default:''" json:"createBy"` // 创建者
	UpdatedBy     string         `gorm:"size:64;default:''" json:"updateBy"` // 更新者
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Users         []User         `gorm:"many2many:user_roles"`
	IsInternal    bool
	OauthClientId *uuid.UUID
	Description   string `gorm:"size:256"`
}

func (Role) TableName() string {
	return "role"
}

func (r *Role) Save() {
	DB().Save(r)
}

func (r *Role) Delete() {
	DB().Delete(r)
}
