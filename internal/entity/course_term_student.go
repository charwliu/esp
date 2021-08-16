package entity

import (
	"time"
)

// CourseTermStudent 培训班级学员表
type CourseTermStudent struct {
	TermID    int64     `gorm:"primary_key;not null" json:"termId"` // 班级ID
	UserID    int64     `gorm:"primary_key;not null" json:"userId"` // 用户ID
	Status    string    `gorm:"size:1;default:0" json:"status"`     // 学员状态（0正常 1停用）
	CreatedAt time.Time `json:"created_at"`                         // 创建时间
	UpdatedAt time.Time `json:"updated_at"`                         // 更新时间
	Term      *CourseTerm
	User      *User
}
