package entity

import (
	"time"
)

// Post 岗位信息表
type Post struct {
	PostID    int64     `gorm:"primary_key;;not null" json:"postId"` // 岗位ID
	PostCode  string    `gorm:"size:64;not null" json:"postCode"`    // 岗位编码
	PostName  string    `gorm:"size:50;not null" json:"postName"`    // 岗位名称
	PostSort  int       `gorm:"not null" json:"postSort"`            // 显示顺序
	Status    string    `gorm:"size:1;not null" json:"status"`       // 状态（0正常 1停用）
	CreatedBy string    `gorm:"size:64;default:''" json:"createBy"`  // 创建者
	UpdatedBy string    `gorm:"size:64;default:''" json:"updateBy"`  // 更新者
	CreatedAt time.Time `json:"created_at"`                          // 创建时间
	UpdatedAt time.Time `json:"updated_at"`                          // 更新时间
	Remark    string    `gorm:"size:500;default:''" json:"remark"`   // 备注信息
}

// PostColumns get sql column name.获取数据库列名
var PostColumns = struct {
	PostID     string
	PostCode   string
	PostName   string
	PostSort   string
	Status     string
	CreateBy   string
	CreateTime string
	UpdateBy   string
	UpdateTime string
	Remark     string
}{
	PostID:     "post_id",
	PostCode:   "post_code",
	PostName:   "post_name",
	PostSort:   "post_sort",
	Status:     "status",
	CreateBy:   "create_by",
	CreateTime: "create_time",
	UpdateBy:   "update_by",
	UpdateTime: "update_time",
	Remark:     "remark",
}
