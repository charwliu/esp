package entity

import (
	"time"
)

// CourseTheme 主题表
type CourseTheme struct {
	ID          int64     `gorm:"primary_key;;not null" json:"themeId"` // 主题编号
	ThemeName   string    `gorm:"size:100" json:"themeName"`            // 主题名称
	ThemeDesc   string    `gorm:"size:1000" json:"themeDesc"`           // 主题描述
	ThemeNum    int       `json:"themeNum"`                             // 显示顺序
	ThemeType   string    `gorm:"size:1" json:"themeType"`              // 主题类型（1培训主题 2资料主题）
	ThemeStatus string    `gorm:"size:1" json:"themeStatus"`            // 主题状态 （0正常 1停用）
	CreatedBy   string    `gorm:"size:64;default:''" json:"createBy"`   // 创建者
	UpdatedBy   string    `gorm:"size:64;default:''" json:"updateBy"`   // 更新者
	CreatedAt   time.Time `json:"created_at"`                           // 创建时间
	UpdatedAt   time.Time `json:"updated_at"`                           // 更新时间
}

// CourseThemeColumns get sql column name.获取数据库列名
var CourseThemeColumns = struct {
	ThemeID     string
	ThemeName   string
	ThemeDesc   string
	ThemeNum    string
	ThemeType   string
	ThemeStatus string
	CreateBy    string
	CreateTime  string
	UpdateBy    string
	UpdateTime  string
}{
	ThemeID:     "theme_id",
	ThemeName:   "theme_name",
	ThemeDesc:   "theme_desc",
	ThemeNum:    "theme_num",
	ThemeType:   "theme_type",
	ThemeStatus: "theme_status",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
}
