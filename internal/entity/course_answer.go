package entity

import (
	"time"
)

// CourseAnswer 答题记录表
type CourseAnswer struct {
	ID         int64     `gorm:"primaryKey;not null" json:"id"` // 编号
	ProgressID int64     `json:"progressId"`                    // 进度编号
	RelationID int64     `json:"relationId"`                    // 关联表编号
	CustomerID int64     `json:"customerId"`                    // 用户Id
	Answer     string    `gorm:"size:200" json:"answer"`        // 答案
	RightNum   int       `gorm:"default:0" json:"rightNum"`     // 答对数
	CreatedAt  time.Time `json:"created_at"`
	Progress   *CourseProgress
	Customer   *User
}

func (CourseAnswer) TableName() string {
	return "course_answer"
}
