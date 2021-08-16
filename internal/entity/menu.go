package entity

import (
	"time"
)

// Menu 菜单权限表
type Menu struct {
	ID        int64     `gorm:"primary_key;not null" json:"menuId"` // 菜单ID
	MenuName  string    `gorm:"size:50;not null" json:"menuName"`   // 菜单名称
	ParentID  int64     `gorm:"default:0" json:"parentId"`          // 父菜单ID
	OrderNum  int       `gorm:"default:0" json:"orderNum"`          // 显示顺序
	URL       string    `gorm:"size:200;default:#" json:"url"`      // 请求地址
	Target    string    `gorm:"size:20;default:''" json:"target"`   // 打开方式（menuItem页签 menuBlank新窗口）
	MenuType  string    `gorm:"size:1;default:''" json:"menuType"`  // 菜单类型（M目录 C菜单 F按钮）
	Visible   string    `gorm:"size:1;default:0" json:"visible"`    // 菜单状态（0显示 1隐藏）
	Perms     string    `gorm:"size:100" json:"perms"`              // 权限标识
	Icon      string    `gorm:"size:100;default:#" json:"icon"`     // 菜单图标
	CreatedBy string    `gorm:"size:64;default:''" json:"createBy"` // 创建者
	UpdatedBy string    `gorm:"size:64;default:''" json:"updateBy"` // 更新者
	CreatedAt time.Time `json:"created_at"`                         // 创建时间
	UpdatedAt time.Time `json:"updated_at"`                         // 更新时间
	Remark    string    `gorm:"size:500;default:''" json:"remark"`  // 备注
	Parent    *Menu
	Roles     []Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;many2many:role_menus;"`
}


// TableName get sql table name.
//  获取数据库表名
func (Menu) TableName() string {
	return "menu"
}
