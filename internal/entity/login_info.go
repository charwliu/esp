package entity

import (
	"time"
)

// LoginInfo 系统访问记录
type LoginInfo struct {
	ID            int64     `gorm:"primary_key;not null" json:"-"`                // 访问ID
	LoginName     string    `gorm:"size:50;default:''" json:"loginName"`          // 登录账号
	IpAddr        string    `gorm:"size:50;default:''" json:"ipaddr"`             // 登录IP地址
	LoginLocation string    `gorm:"size:255;default:''" json:"loginLocation"`     // 登录地点
	Browser       string    `gorm:"size:50;default:''" json:"browser"`            // 浏览器类型
	Os            string    `gorm:"size:50;default:''" json:"os"`                 // 操作系统
	Status        string    `gorm:"size:1;default:0" json:"status"`               // 登录状态（0成功 1失败）
	Msg           string    `gorm:"size:255;default:''" json:"msg"`               // 提示消息
	LoginAt       time.Time `gorm:"column:login_at;type:datetime" json:"loginAt"` // 访问时间
}
