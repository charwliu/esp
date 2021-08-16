package entity

import (
	"time"
)

// CourseArticleBak [...]
type CourseArticleBak struct {
	ArticleID      int64     `gorm:"column:article_id;type:bigint(20);not null;default:0" json:"articleId"` // 编号
	ClassifyID     string    `gorm:"column:classify_id;type:varchar(32)" json:"classifyId"`                 // 文章所属分类
	ArticleTitle   string    `gorm:"column:article_title;type:varchar(100)" json:"articleTitle"`            // 标题
	ArticleContent string    `gorm:"column:article_content;type:longtext" json:"articleContent"`            // 内容
	ArticleTime    *int      `gorm:"column:article_time;type:int(11)" json:"articleTime"`                   // 阅读时间
	WordNumber     *int      `gorm:"column:word_number;type:int(11)" json:"wordNumber"`                     // 文章字数
	ArticleLabel   string    `gorm:"column:article_label;type:varchar(100)" json:"articleLabel"`            // 文章标签
	ArticleStatus  string    `gorm:"column:article_status;type:char(1)" json:"articleStatus"`               // 状态 （0正常 1停用）
	CreatedBy      string    `gorm:"size:64;default:''" json:"createBy"`                                    // 创建者
	UpdatedBy      string    `gorm:"size:64;default:''" json:"updateBy"`                                    // 更新者
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// CourseArticleBakColumns get sql column name.获取数据库列名
var CourseArticleBakColumns = struct {
	ArticleID      string
	ClassifyID     string
	ArticleTitle   string
	ArticleContent string
	ArticleTime    string
	WordNumber     string
	ArticleLabel   string
	ArticleStatus  string
	CreateBy       string
	CreateTime     string
	UpdateBy       string
	UpdateTime     string
}{
	ArticleID:      "article_id",
	ClassifyID:     "classify_id",
	ArticleTitle:   "article_title",
	ArticleContent: "article_content",
	ArticleTime:    "article_time",
	WordNumber:     "word_number",
	ArticleLabel:   "article_label",
	ArticleStatus:  "article_status",
	CreateBy:       "create_by",
	CreateTime:     "create_time",
	UpdateBy:       "update_by",
	UpdateTime:     "update_time",
}
