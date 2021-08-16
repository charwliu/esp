package entity

// CourseRelation 课程内容关联表
type CourseRelation struct {
	ID           int64  `gorm:"primary_ey;not null" json:"id"` // 编号
	CourseID     int64  `json:"courseId"`                      // 课程编号
	CategoryID   string `gorm:"size:32" json:"categoryId"`     // 内容分类编号
	QuestionType string `gorm:"size:32" json:"questionType"`   // 题型
	QuestionID   int64  `json:"questionId"`                    // 题号
	QuestionSort int    `json:"questionSort"`                  // 排序
	Course       *CourseInfo
}
