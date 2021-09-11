package entity

import (
	"time"
)

// CourseChoice 选择题表
type CourseChoice struct {
	ID             int64     `gorm:"primaryKey;not null" json:"Id"`      // 选择题编号
	ChoiceQuestion string    `gorm:"type:text" json:"choiceQuestion"`    // 题目
	OptionA        string    `gorm:"size:200" json:"optionA"`            // 选项A
	OptionB        string    `gorm:"size:200" json:"optionB"`            // 选项B
	OptionC        string    `gorm:"size:200" json:"optionC"`            // 选项C
	OptionD        string    `gorm:"size:200" json:"optionD"`            // 选项D
	OptionE        string    `gorm:"size:200" json:"optionE"`            // 选项E
	OptionF        string    `gorm:"size:200" json:"optionF"`            // 选项F
	OptionG        string    `gorm:"size:200" json:"optionG"`            // 选项G
	OptionH        string    `gorm:"size:200" json:"optionH"`            // 选项H
	OptionI        string    `gorm:"size:200" json:"optionI"`            // 选项I
	ChoiceAnswer   string    `gorm:"size:32" json:"choiceAnswer"`        // 答案
	ChoiceScore    *int8     `gorm:"size:1" json:"choiceScore"`          // 分值
	ChoiceExplain  string    `gorm:"type:text" json:"choiceExplain"`     // 详解
	SupportAnswer  string    `gorm:"size:200" json:"supportAnswer"`      // 支持答案
	ChoiceStatus   string    `gorm:"size:1" json:"choiceStatus"`         // 状态 （0正常 1停用）
	ChoiceLabel    string    `gorm:"size:10" json:"choiceLabel"`         // 题目标签
	CreatedBy      string    `gorm:"size:64;default:''" json:"createBy"` // 创建者
	UpdatedBy      string    `gorm:"size:64;default:''" json:"updateBy"` // 更新者
	CreatedAt      time.Time `json:"created_at"`                         // 创建时间
	UpdatedAt      time.Time `json:"updated_at"`                         // 更新时间
}

func (d CourseChoice) TableName() string {
	return "course_choice"
}
