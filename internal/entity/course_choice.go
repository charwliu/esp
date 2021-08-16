package entity

import (
	"time"
)

// CourseChoice 选择题表
type CourseChoice struct {
	ChoiceID       int64     `gorm:"primaryKey;column:choice_id;type:bigint(20);not null" json:"choiceId"` // 选择题编号
	ChoiceQuestion string    `gorm:"column:choice_question;type:text" json:"choiceQuestion"`               // 题目
	OptionA        string    `gorm:"column:optionA;type:varchar(200)" json:"optionA"`                      // 选项A
	OptionB        string    `gorm:"column:optionB;type:varchar(200)" json:"optionB"`                      // 选项B
	OptionC        string    `gorm:"column:optionC;type:varchar(200)" json:"optionC"`                      // 选项C
	OptionD        string    `gorm:"column:optionD;type:varchar(200)" json:"optionD"`                      // 选项D
	OptionE        string    `gorm:"column:optionE;type:varchar(200)" json:"optionE"`                      // 选项E
	OptionF        string    `gorm:"column:optionF;type:varchar(200)" json:"optionF"`                      // 选项F
	OptionG        string    `gorm:"column:optionG;type:varchar(200)" json:"optionG"`                      // 选项G
	OptionH        string    `gorm:"column:optionH;type:varchar(200)" json:"optionH"`                      // 选项H
	OptionI        string    `gorm:"column:optionI;type:varchar(200)" json:"optionI"`                      // 选项I
	ChoiceAnswer   string    `gorm:"column:choice_answer;type:varchar(32)" json:"choiceAnswer"`            // 答案
	ChoiceScore    *int8     `gorm:"column:choice_score;type:tinyint(4)" json:"choiceScore"`               // 分值
	ChoiceExplain  string    `gorm:"column:choice_explain;type:text" json:"choiceExplain"`                 // 详解
	SupportAnswer  string    `gorm:"column:support_answer;type:varchar(200)" json:"supportAnswer"`         // 支持答案
	ChoiceStatus   string    `gorm:"column:choice_status;type:char(1)" json:"choiceStatus"`                // 状态 （0正常 1停用）
	ChoiceLabel    string    `gorm:"column:choice_label;type:varchar(10)" json:"choiceLabel"`              // 题目标签
	CreatedBy      string    `gorm:"size:64;default:''" json:"createBy"`                                   // 创建者
	UpdatedBy      string    `gorm:"size:64;default:''" json:"updateBy"`                                   // 更新者
	CreatedAt      time.Time `json:"created_at"`                                                           // 创建时间
	UpdatedAt      time.Time `json:"updated_at"`                                                           // 更新时间
}

// CourseChoiceColumns get sql column name.获取数据库列名
var CourseChoiceColumns = struct {
	ChoiceID       string
	ChoiceQuestion string
	OptionA        string
	OptionB        string
	OptionC        string
	OptionD        string
	OptionE        string
	OptionF        string
	OptionG        string
	OptionH        string
	OptionI        string
	ChoiceAnswer   string
	ChoiceScore    string
	ChoiceExplain  string
	SupportAnswer  string
	ChoiceStatus   string
	ChoiceLabel    string
	CreateBy       string
	CreateTime     string
	UpdateBy       string
	UpdateTime     string
}{
	ChoiceID:       "choice_id",
	ChoiceQuestion: "choice_question",
	OptionA:        "optionA",
	OptionB:        "optionB",
	OptionC:        "optionC",
	OptionD:        "optionD",
	OptionE:        "optionE",
	OptionF:        "optionF",
	OptionG:        "optionG",
	OptionH:        "optionH",
	OptionI:        "optionI",
	ChoiceAnswer:   "choice_answer",
	ChoiceScore:    "choice_score",
	ChoiceExplain:  "choice_explain",
	SupportAnswer:  "support_answer",
	ChoiceStatus:   "choice_status",
	ChoiceLabel:    "choice_label",
	CreateBy:       "create_by",
	CreateTime:     "create_time",
	UpdateBy:       "update_by",
	UpdateTime:     "update_time",
}
