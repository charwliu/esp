package entity

import (
	"time"
)

// CourseTerm 学期表
type CourseTerm struct {
	ID         int64     `gorm:"primary_key;;not null" json:"termId"` // 学期编号
	DeptID     int64     `gorm:"" json:"deptId"`                      // 部门
	TermNum    int       `json:"termNum"`                             // 学期数
	TermTitle  string    `gorm:"size:100" json:"termTitle"`           // 学期标题
	BeginDate  time.Time `json:"beginDate"`                           // 开课时间
	EndDate    time.Time `json:"endDate"`                             // 截止时间
	ClassHour  string    `gorm:"size:32" json:"classHour"`            // 学期课时
	TermPrice  float64   `gorm:"type:decimal(32,2)" json:"termPrice"` // 价格
	TermStatus string    `gorm:"size:1" json:"termStatus"`            // 状态 （0正常 1停用）
	CreatedBy  string    `gorm:"size:64;default:''" json:"createBy"`  // 创建者
	UpdatedBy  string    `gorm:"size:64;default:''" json:"updateBy"`  // 更新者
	CreatedAt  time.Time `json:"created_at"`                          // 创建时间
	UpdatedAt  time.Time `json:"updated_at"`                          // 更新时间
	Dept       *Dept
}

// CourseTermColumns get sql column name.获取数据库列名
var CourseTermColumns = struct {
	TermID     string
	DeptID     string
	TermNum    string
	TermTitle  string
	BeginDate  string
	EndDate    string
	ClassHour  string
	TermPrice  string
	TermStatus string
	CreateBy   string
	CreateTime string
	UpdateBy   string
	UpdateTime string
}{
	TermID:     "term_id",
	DeptID:     "dept_id",
	TermNum:    "term_num",
	TermTitle:  "term_title",
	BeginDate:  "begin_date",
	EndDate:    "end_date",
	ClassHour:  "class_hour",
	TermPrice:  "term_price",
	TermStatus: "term_status",
	CreateBy:   "create_by",
	CreateTime: "create_time",
	UpdateBy:   "update_by",
	UpdateTime: "update_time",
}
