package entity

import (
	"time"
)

// Position 岗位信息表
type Position struct {
	ID        int64     `gorm:"primary_key;;not null" json:"Id"`   // 岗位ID
	Code      string    `gorm:"size:64;not null" json:"Code"`      // 岗位编码
	Name      string    `gorm:"size:50;not null" json:"Name"`      // 岗位名称
	Sort      int       `gorm:"not null" json:"Sort"`              // 显示顺序
	Status    string    `gorm:"size:1;not null" json:"Status"`     // 状态（0正常 1停用）
	CreatedBy string    `gorm:"size:64;default:''" json:"-"`       // 创建者
	UpdatedBy string    `gorm:"size:64;default:''" json:"-"`       // 更新者
	CreatedAt time.Time `json:"createdAt"`                         // 创建时间
	UpdatedAt time.Time `json:"updatedAt"`                         // 更新时间
	Remark    string    `gorm:"size:500;default:''" json:"Remark"` // 备注信息
}

func (Position) TableName() string {
	return "position"
}


func (m *Position) Save() error {
	return DB().Save(m).Error
}

func (m *Position) Create() error {
	return DB().Create(m).Error
}

func (m *Position) Delete() error {
	return DB().Delete(m).Error
}

