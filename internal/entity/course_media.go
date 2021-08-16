package entity

import (
	"time"
)

// CourseMedia 资料表
type CourseMedia struct {
	ID          int64     `gorm:"primary_key;not null" json:"mediaId"` // 资料编号
	ThemeID     int64     `json:"themeId"`                             // 主题编号
	MediaName   string    `gorm:"size:100" json:"mediaName"`           // 资料标题
	MediaDesc   string    `gorm:"size:2000" json:"mediaDesc"`          // 资料描述
	MediaLink   string    `gorm:"size:1000" json:"mediaLink"`          // 链接地址
	MediaImage  string    `gorm:"size:1000" json:"mediaImage"`         // 图片路径
	MediaNum    int       `json:"mediaNum"`                            // 显示顺序
	MediaStatus string    `gorm:"size:1" json:"mediaStatus"`           // 资料状态 （0正常 1停用）
	CreatedBy   string    `gorm:"size:64;default:''" json:"createBy"`  // 创建者
	UpdatedBy   string    `gorm:"size:64;default:''" json:"updateBy"`  // 更新者
	CreatedAt   time.Time `json:"created_at"`                          // 创建时间
	UpdatedAt   time.Time `json:"updated_at"`                          // 更新时间
	Theme       *CourseTheme
}
