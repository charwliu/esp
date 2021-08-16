package entity

import (
	"time"
)

// CourseRead 文章阅读记录表
type CourseRead struct {
	ID         int64     `gorm:"primary_key;not null" json:"id"` // 编号
	ProgressID int64     `json:"progressId"`                     // 进度编号
	CourseID   int64     `json:"courseId"`                       // 课程编号
	CategoryID string    `json:"categoryId"`                     // 内容分类编号
	CustomerID int64     `json:"customerId"`                     // 用户Id
	BeginTime  time.Time `json:"beginTime"`                      // 开始时间
	EndTime    time.Time `json:"endTime"`                        // 结束时间
	ReadFlag   string    `gorm:"size:1" json:"readFlag"`         // 阅读标识（1正常阅读 2回看阅读）
	Progress   *CourseProgress
	Course     *CourseDesign
	Customer   *User
}
