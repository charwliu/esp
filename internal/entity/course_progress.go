package entity

import (
	"time"
)

// CourseProgress 课程进度表
type CourseProgress struct {
	ID           int64     `gorm:"primary_key;not null" json:"id"`     // 编号
	CourseID     int64     `json:"courseId"`                           // 课程编号
	BeginTime    time.Time `json:"beginTime"`                          // 开始时间
	CustomerID   int64     `json:"customerId"`                         // 用户Id
	PlaybackNum  int       `json:"playbackNum"`                        // 回看次数
	PlaybackTime float64   `json:"playbackTime"`                       // 回看时长
	ReadTime     float64   `json:"readTime"`                           // 阅读考试用时
	AnswerScore  int       `json:"answerScore"`                        // 分数
	ExamTotal    int       `json:"examTotal"`                          // 题目总数
	AnswerRight  int       `json:"answerRight"`                        // 题目答对数
	AnswerRate   string    `gorm:"size:10" json:"answerRate"`          // 答题正确率
	CourseStatus string    `gorm:"size:1" json:"courseStatus"`         // 状态 （0学习中 1已学习 2补学习）
	CreatedBy    string    `gorm:"size:64;default:''" json:"createBy"` // 创建者
	CreatedAt    time.Time `json:"created_at"`                         // 创建时间
	Course       *CourseInfo
	Customer     *User
}

func (CourseProgress) TableName() string {
	return "course_progress"
}
