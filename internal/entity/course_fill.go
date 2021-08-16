package entity

import (
	"time"
)

// CourseFill 填空题表
type CourseFill struct {
	FillID        int64     `gorm:"primaryKey;column:fill_id;type:bigint(20);not null" json:"fillId"` // 填空题编号
	FillQuestion  string    `gorm:"column:fill_question;type:text" json:"fillQuestion"`               // 题目
	FillAnswer    string    `gorm:"column:fill_answer;type:text" json:"fillAnswer"`                   // 答案
	FillScore     *int8     `gorm:"column:fill_score;type:tinyint(4)" json:"fillScore"`               // 分值
	FillExplain   string    `gorm:"column:fill_explain;type:text" json:"fillExplain"`                 // 详解
	SupportAnswer string    `gorm:"column:support_answer;type:varchar(200)" json:"supportAnswer"`     // 支持答案
	FillStatus    string    `gorm:"column:fill_status;type:char(1)" json:"fillStatus"`                // 状态 （0正常 1停用）
	FillLabel     string    `gorm:"column:fill_label;type:varchar(10)" json:"fillLabel"`              // 题目标签
	CreatedBy     string    `gorm:"size:64;default:''" json:"createBy"`                               // 创建者
	UpdatedBy     string    `gorm:"size:64;default:''" json:"updateBy"`                               // 更新者
	CreatedAt     time.Time `json:"created_at"`                                                       // 创建时间
	UpdatedAt     time.Time `json:"updated_at"`                                                       // 更新时间
}

// CourseFillColumns get sql column name.获取数据库列名
var CourseFillColumns = struct {
	FillID        string
	FillQuestion  string
	FillAnswer    string
	FillScore     string
	FillExplain   string
	SupportAnswer string
	FillStatus    string
	FillLabel     string
	CreateBy      string
	CreateTime    string
	UpdateBy      string
	UpdateTime    string
}{
	FillID:        "fill_id",
	FillQuestion:  "fill_question",
	FillAnswer:    "fill_answer",
	FillScore:     "fill_score",
	FillExplain:   "fill_explain",
	SupportAnswer: "support_answer",
	FillStatus:    "fill_status",
	FillLabel:     "fill_label",
	CreateBy:      "create_by",
	CreateTime:    "create_time",
	UpdateBy:      "update_by",
	UpdateTime:    "update_time",
}
