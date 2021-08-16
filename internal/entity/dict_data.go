package entity

import (
	"time"
)

// DictData 字典数据表
type DictData struct {
	DictCode  int64     `gorm:"primary_key;column:dict_code;not null" json:"dictCode"` // 字典编码
	DictSort  int       `gorm:"default:0" json:"dictSort"`                             // 字典排序
	DictLabel string    `gorm:"size:100;default:''" json:"dictLabel"`                  // 字典标签
	DictValue string    `gorm:"size:100;default:''" json:"dictValue"`                  // 字典键值
	DictType  string    `gorm:"size:100;default:''" json:"dictType"`                   // 字典类型
	CSSClass  string    `gorm:"size:100" json:"cssClass"`                              // 样式属性（其他样式扩展）
	ListClass string    `gorm:"size:100" json:"listClass"`                             // 表格回显样式
	IsDefault string    `gorm:"size:1;default:N" json:"isDefault"`                     // 是否默认（Y是 N否）
	Status    string    `gorm:"size:1;default:0" json:"status"`                        // 状态（0正常 1停用）
	CreatedBy string    `gorm:"size:64;default:''" json:"createBy"`                    // 创建者
	UpdatedBy string    `gorm:"size:64;default:''" json:"updateBy"`                    // 更新者
	CreatedAt time.Time `json:"created_at"`                                            // 创建时间
	UpdatedAt time.Time `json:"updated_at"`                                            // 更新时间
	Remark    string    `gorm:"size:500;default:''" json:"remark"`                     // 备注信息
}
