package entity

import (
	"time"
)

// UserLocation 用户签到记录表
type UserLocation struct {
	ID        int64     `gorm:"primary_key;not null" json:"id"` // 主键ID
	UserID    int64     `json:"userId"`                         // 用户ID
	TermID    int64     `json:"termId"`                         // 培训班编号
	Latitude  string    `gorm:"size:50" json:"latitude"`        // 纬度
	Longitude string    `gorm:"size:50" json:"longitude"`       // 经度
	Precision string    `gorm:"size:50" json:"precision"`       // 精度
	Province  string    `gorm:"size:100" json:"province"`       // 省
	City      string    `gorm:"size:100" json:"city"`           // 市
	Area      string    `gorm:"size:200" json:"area"`           // 地区
	State     string    `gorm:"size:1" json:"state"`            // 状态（0允许 1拒绝）
	CreatedAt time.Time `json:"created_at"`                     // 创建时间
	User      *User
	Term      *CourseTerm
}


func (UserLocation) TableName() string {
	return "user_location"
}
