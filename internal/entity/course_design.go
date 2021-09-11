package entity

import (
	"time"
)

// CourseDesign 设计题表
type CourseDesign struct {
	ID             int64     `gorm:"primary_key;not null" json:"designId"` // 设计题编号
	DesignQuestion string    `gorm:"type:text" json:"designQuestion"`      // 题目
	DesignAnswer   string    `gorm:"type:text" json:"designAnswer"`        // 答案
	DesignExplain  string    `gorm:"type:text" json:"designExplain"`       // 详解
	DesignStatus   string    `gorm:"size:1" json:"designStatus"`           // 状态 （0正常 1停用）
	DesignLabel    string    `gorm:"size:10" json:"designLabel"`           // 题目标签
	CreatedBy      string    `gorm:"size:64;default:''" json:"createBy"`   // 创建者
	UpdatedBy      string    `gorm:"size:64;default:''" json:"updateBy"`   // 更新者
	CreatedAt      time.Time `json:"created_at"`                           // 创建时间
	UpdatedAt      time.Time `json:"updated_at"`                           // 更新时间
}

func (d CourseDesign) TableName() string  {
	return "course_design"
}

