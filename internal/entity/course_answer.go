package entity

import (
	"time"
)

// CourseAnswer 答题记录表
type CourseAnswer struct {
	ID         int64     `gorm:"primaryKey;column:id;type:bigint(20);not null" json:"id"`          // 编号
	ProgressID int64     `json:"progressId"`                                                       // 进度编号
	RelationID int64     `json:"relationId"`                                                       // 关联表编号
	CustomerID int64     `json:"customerId"`                                                       // 用户Id
	Answer     string    `gorm:"column:answer;type:varchar(200)" json:"answer"`                    // 答案
	RightNum   int       `gorm:"column:right_num;type:int(11);not null;default:0" json:"rightNum"` // 答对数
	CreatedAt  time.Time `json:"created_at"`
	Progress   *CourseProgress
	Customer   *User
}

// CourseAnswerColumns get sql column name.获取数据库列名
var CourseAnswerColumns = struct {
	ID         string
	ProgressID string
	RelationID string
	CustomerID string
	Answer     string
	RightNum   string
	CreateTime string
}{
	ID:         "id",
	ProgressID: "progress_id",
	RelationID: "relation_id",
	CustomerID: "customer_id",
	Answer:     "answer",
	RightNum:   "right_num",
	CreateTime: "create_time",
}
