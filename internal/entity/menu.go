package entity

import (
	"time"

	"go.uber.org/zap"
)

// Menu 菜单权限表
type Menu struct {
	ID        int64     `gorm:"primary_key;not null" json:"menuId"` // 菜单ID
	Name      string    `gorm:"size:50;index;not null" json:"Name"` // 菜单名称
	OrderNum  int       `gorm:"default:0" json:"OrderNum"`          // 显示顺序
	URL       string    `gorm:"size:200;default:#" json:"Url"`      // 请求地址
	Target    string    `gorm:"size:20;default:''" json:"Target"`   // 打开方式（menuItem页签 menuBlank新窗口）
	MenuType  string    `gorm:"size:1;default:''" json:"MenuType"`  // 菜单类型（M目录 C菜单 F按钮）
	Visible   string    `gorm:"size:1;default:0" json:"Visible"`    // 菜单状态（0显示 1隐藏）
	Perms     string    `gorm:"size:100" json:"Perms"`              // 权限标识
	Icon      string    `gorm:"size:100;default:#" json:"Icon"`     // 菜单图标
	CreatedBy string    `gorm:"size:64;default:''" json:"-"`        // 创建者
	UpdatedBy string    `gorm:"size:64;default:''" json:"-"`        // 更新者
	CreatedAt time.Time `json:"createdAt"`                          // 创建时间
	UpdatedAt time.Time `json:"updatedAt"`                          // 更新时间
	Remark    string    `gorm:"size:500;default:''" json:"Remark"`  // 备注
	ParentID  *int64    `json:"ParentId"`                           // 父菜单ID
	SubMenu   []Menu    `gorm:"foreignKey:ParentID" json:"SubMenu"`
	Roles     []Role    `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;many2many:role_menu;"`
}

// TableName get sql table name.
func (Menu) TableName() string {
	return "menu"
}

func (m *Menu) Save() error {
	return DB().Save(m).Error
}

func (m *Menu) Create() error {
	return DB().Create(m).Error
}

func (m *Menu) Delete() error {
	return DB().Delete(m).Error
}

// UnknownMenu is PhotoPrism's default Menu.
var UnknownMenu = Menu{
	ID:   0,
	Name: "Unknown",
}

// CreateUnknownMenu creates the default Menu if not exists.
func CreateUnknownMenu() {
	FirstOrCreateMenu(&UnknownMenu)
}

func FirstOrCreateMenu(m *Menu) *Menu {
	if m.Name == "" {
		L().Error("menu: menu must not be empty (find or create)")
		return nil
	}

	result := Menu{}

	if findErr := DB().Where("id = ? OR name = ?", m.ID, m.Name).First(&result).Error; findErr == nil {
		return &result
	} else if createErr := m.Create(); createErr == nil {
		return m
	} else if err := DB().Where("id = ? OR name = ?", m.ID, m.Name).First(&result).Error; err == nil {
		return &result
	} else {
		L().Error("menu: create menu",
			zap.Int64("id", m.ID),
			zap.String("name", m.Name),
			zap.Error(createErr))
	}
	return &result
}

func FindMenuByName(name string) *Menu {
	if name == "" {
		L().Info("name is empty")
		return nil
	}
	menu := Menu{}

	if err := DB().Preload("SubMenu").Where("name = ?", name).First(&menu).Error; err == nil {
		return &menu
	} else {
		L().Info("menu not found", zap.String("name", name))
		return nil
	}
}
