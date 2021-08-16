package entity

import (
	"time"
)

// CourseArticle 阅读文章表
type CourseArticle struct {
	ID             int64     `gorm:"primary_key;not null" json:"articleId"` // 编号
	ClassifyID     string    `gorm:"size:32" json:"classifyId"`             // 文章所属分类
	ArticleTitle   string    `gorm:"size:100" json:"articleTitle"`          // 标题
	ArticleContent string    `gorm:"type:longtext" json:"articleContent"`   // 内容
	ArticleTime    int       `json:"articleTime"`                           // 阅读时间
	WordNumber     int       `json:"wordNumber"`                            // 文章字数
	ArticleLabel   string    `gorm:"size:100" json:"articleLabel"`          // 文章标签
	ArticleStatus  string    `gorm:"size:1" json:"articleStatus"`           // 状态 （0正常 1停用）
	CreatedBy      string    `gorm:"size:64;default:''" json:"createBy"`    // 创建者
	UpdatedBy      string    `gorm:"size:64;default:''" json:"updateBy"`    // 更新者
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// CourseArticleColumns get sql column name.获取数据库列名
var CourseArticleColumns = struct {
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
