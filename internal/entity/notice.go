package entity

import (
	"time"
)

// Notice 通知公告表
type Notice struct {
	ID            int       `gorm:"primary_key;not null" json:"noticeId"` // 公告ID
	NoticeTitle   string    `gorm:"size:50;not null" json:"noticeTitle"`  // 公告标题
	NoticeType    string    `gorm:"size:1;not null" json:"noticeType"`    // 公告类型（1通知 2公告）
	NoticeContent string    `gorm:"size:2000" json:"noticeContent"`       // 公告内容
	Status        string    `gorm:"size:1;default:0" json:"status"`       // 公告状态（0正常 1关闭）
	CreatedBy     string    `gorm:"size:64;default:''" json:"createBy"`   // 创建者
	UpdatedBy     string    `gorm:"size:64;default:''" json:"updateBy"`   // 更新者
	CreatedAt     time.Time `json:"created_at"`                           // 创建时间
	UpdatedAt     time.Time `json:"updated_at"`                           // 更新时间
	Remark        string    `gorm:"size:500;default:''" json:"remark"`    // 备注
}
