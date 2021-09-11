package entity

import (
	"time"
)

// CourseJudge 判断题表
type CourseJudge struct {
	ID            int64     `gorm:"primary_key;not null" json:"judgeId"` // 判断题编号
	JudgeQuestion string    `gorm:"type:text" json:"judgeQuestion"`      // 题目
	JudgeAnswer   string    `gorm:"type:text" json:"judgeAnswer"`        // 答案
	JudgeScore    int8      `json:"judgeScore"`                          // 分值
	JudgeExplain  string    `gorm:"type:text" json:"judgeExplain"`       // 详解
	JudgeStatus   string    `gorm:"size:1" json:"judgeStatus"`           // 状态 （0正常 1停用）
	JudgeLabel    string    `gorm:"size:10" json:"judgeLabel"`           // 题目标签
	CreatedBy     string    `gorm:"size:64;default:''" json:"createBy"`  // 创建者
	UpdatedBy     string    `gorm:"size:64;default:''" json:"updateBy"`  // 更新者
	CreatedAt     time.Time `json:"created_at"`                          // 创建时间
	UpdatedAt     time.Time `json:"updated_at"`                          // 更新时间
}

func (d CourseJudge) TableName() string {
	return "course_judge"
}