package entity

import (
	"time"
)

// UserOnline 在线用户记录
type UserOnline struct {
	SessionID      string    `gorm:"primary_ey;size:50;not null;default:''" json:"sessionId"` // 用户会话id
	LoginName      string    `gorm:"size:50;default:''" json:"loginName"`                     // 登录账号
	DeptName       string    `gorm:"size:50;default:''" json:"deptName"`                      // 部门名称
	IpAddr         string    `gorm:"size:50;default:''" json:"ipaddr"`                        // 登录IP地址
	LoginLocation  string    `gorm:"size:255;default:''" json:"loginLocation"`                // 登录地点
	Browser        string    `gorm:"size:50;default:''" json:"browser"`                       // 浏览器类型
	Os             string    `gorm:"size:50;default:''" json:"os"`                            // 操作系统
	Status         string    `gorm:"size:10;default:''" json:"status"`                        // 在线状态on_line在线off_line>离线
	StartedAt      time.Time `json:"StartedAt"`                                               // session创建时间
	LastAccessedAt time.Time `json:"lastAccessedAt"`                                          // session最后访问时间
	ExpireTime     int       `gorm:"default:0" json:"expireTime"`                             // 超时时间，单位为分钟
}

func (UserOnline) TableName() string {
	return "user_online"
}
