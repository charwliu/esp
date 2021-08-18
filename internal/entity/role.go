package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Roles []*Role

// Role 角色信息表
type Role struct {
	ID            int64          `gorm:"primary_key;not null" json:"roleId"`          // 角色ID
	RoleName      string         `gorm:"size:32;not null" json:"roleName,omitempty"`  // 角色名称
	RoleKey       string         `gorm:"size:100;not null" json:"roleKey,omitempty"`  // 角色权限字符串
	RoleSort      int            `gorm:"not null" json:"roleSort,omitempty"`          // 显示顺序
	DataScope     string         `gorm:"size:1;default:1" json:"dataScope,omitempty"` // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Status        string         `gorm:"size:1;not null" json:"status,omitempty"`     // 角色状态（0正常 1停用）
	UpdateTime    time.Time      `json:"-"`                                           // 更新时间
	Remark        string         `gorm:"size:500" json:"remark,omitempty"`            // 备注
	CreatedBy     string         `gorm:"size:64;default:''" json:"-"`                 // 创建者
	UpdatedBy     string         `gorm:"size:64;default:''" json:"-"`                 // 更新者
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	Users         []User         `gorm:"many2many:user_roles" json:"users,omitempty"`
	Description   string         `gorm:"size:256" json:"description,omitempty"`
	IsInternal    bool           `json:"isInternal,omitempty"`
	OauthClientId *uuid.UUID     `json:"oauth_client_id,omitempty"`
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
