package entity

import (
	"time"
)

// DictType 字典类型表
type DictType struct {
	ID        int64     `gorm:"primary_key;not null" json:"dictId"`         // 字典主键
	DictName  string    `gorm:"size:100;default:''" json:"dictName"`        // 字典名称
	DictType  string    `gorm:"unique;size:100;default:''" json:"dictType"` // 字典类型
	Status    string    `gorm:"size:1;default:0" json:"status"`             // 状态（0正常 1停用）
	CreatedBy string    `gorm:"size:64;default:''" json:"createBy"`         // 创建者
	UpdatedBy string    `gorm:"size:64;default:''" json:"updateBy"`         // 更新者
	CreatedAt time.Time `json:"created_at"`                                 // 创建时间
	UpdatedAt time.Time `json:"updated_at"`                                 // 更新时间
	Remark    string    `gorm:"size:500;default:''" json:"remark"`          // 备注信息
}
