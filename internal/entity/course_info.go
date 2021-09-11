package entity

import (
	"time"
)

// CourseInfo 课程表
type CourseInfo struct {
	ID            int64     `gorm:"primary_key;not null" json:"courseId"`             // 课程编号
	TermID        int64     `json:"termId"`                                           // 部门编号
	CourseNum     int64     `json:"courseNum"`                                        // 主题编号
	CourseName    string    `gorm:"size:100" json:"courseName"`                       // 课程名称
	SubjectID     string    `gorm:"size:32" json:"subjectId"`                         // 科目
	AuditStatus   string    `gorm:"size:1;default:0" json:"auditStatus"`              // 审核状态 （0未审核 1已审核）
	ReleaseTime   time.Time `json:"releaseTime"`                                      // 发布时间
	ReleaseStatus string    `gorm:"size:1;default:0" json:"releaseStatus"`            // 发布状态 （0未发布 1已发布）
	CourseStatus  string    `gorm:"size:1" json:"courseStatus"`                       // 课程状态 （0正常 1停用）
	StickyStatus  string    `gorm:"size:1" json:"stickyStatus"`                       // 置顶状态（N否 Y是）
	PassRate      int       `json:"passRate"`                                         // 及格线
	CreatedBy     string    `gorm:"size:64;default:''" json:"createBy"`               // 创建者
	UpdatedBy     string    `gorm:"size:64;default:''" json:"updateBy"`               // 更新者
	CreatedAt     time.Time `json:"created_at"`                                       // 创建时间
	UpdatedAt     time.Time `json:"updated_at"`                                       // 更新时间
	AuditedBy     string    `gorm:"size:64" json:"auditedBy"`                         // 审核人
	AuditedAt     time.Time `gorm:"column:audited_at;type:datetime" json:"auditedAt"` // 审核时间
	PassingScore  int8      `gorm:"default:0" json:"passingScore"`
	Term          *CourseTerm
}

// CourseInfoColumns get sql column name.获取数据库列名
var CourseInfoColumns = struct {
	CourseID      string
	TermID        string
	CourseNum     string
	CourseName    string
	SubjectID     string
	AuditStatus   string
	ReleaseTime   string
	ReleaseStatus string
	CourseStatus  string
	PassingScore  string
	StickyStatus  string
	PassRate      string
	CreateBy      string
	CreateTime    string
	UpdateBy      string
	UpdateTime    string
	AuditBy       string
	AuditTime     string
}{
	CourseID:      "course_id",
	TermID:        "term_id",
	CourseNum:     "course_num",
	CourseName:    "course_name",
	SubjectID:     "subject_id",
	AuditStatus:   "audit_status",
	ReleaseTime:   "release_time",
	ReleaseStatus: "release_status",
	CourseStatus:  "course_status",
	PassingScore:  "passing_score",
	StickyStatus:  "sticky_status",
	PassRate:      "pass_rate",
	CreateBy:      "create_by",
	CreateTime:    "create_time",
	UpdateBy:      "update_by",
	UpdateTime:    "update_time",
	AuditBy:       "audit_by",
	AuditTime:     "audit_time",
}
