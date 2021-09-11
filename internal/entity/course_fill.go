package entity

import (
	"time"
)

// CourseFill 填空题表
type CourseFill struct {
	ID            int64     `gorm:"primaryKey;not null" json:"fillId"`  // 填空题编号
	FillQuestion  string    `gorm:"type:text" json:"fillQuestion"`      // 题目
	FillAnswer    string    `gorm:"type:text" json:"fillAnswer"`        // 答案
	FillScore     *int8     `gorm:"size:1" json:"fillScore"`            // 分值
	FillExplain   string    `gorm:"type:text" json:"fillExplain"`       // 详解
	SupportAnswer string    `gorm:"size:200" json:"supportAnswer"`      // 支持答案
	FillStatus    string    `gorm:"size:1" json:"fillStatus"`           // 状态 （0正常 1停用）
	FillLabel     string    `gorm:"size:10" json:"fillLabel"`           // 题目标签
	CreatedBy     string    `gorm:"size:64;default:''" json:"createBy"` // 创建者
	UpdatedBy     string    `gorm:"size:64;default:''" json:"updateBy"` // 更新者
	CreatedAt     time.Time `json:"created_at"`                         // 创建时间
	UpdatedAt     time.Time `json:"updated_at"`                         // 更新时间
}

func (d CourseFill) TableName() string {
	return "course_fill"
}
