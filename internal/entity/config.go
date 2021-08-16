package entity

import (
	"time"

	"gorm.io/gorm"
)

// Config 参数配置表
type Config struct {
	ConfigID    int            `gorm:"primaryKey;column:config_id;type:int(5);not null"` // 参数主键
	ConfigName  string         `gorm:"size:100;default:''"`                              // 参数名称
	ConfigKey   string         `gorm:"size:100;default:''"`                              // 参数键名
	ConfigValue string         `gorm:"size:500;default:''"`                              // 参数键值
	ConfigType  string         `gorm:"size:1;default:N"`                                 // 系统内置（Y是 N否）
	CreatedBy   string         `gorm:"size:64;default:''" json:"createBy"`               // 创建者
	UpdatedBy   string         `gorm:"size:64;default:''" json:"updateBy"`               // 更新者
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Remark      string         `gorm:"size:500;default:''" json:"remark"` // 备注信息
}

// ConfigColumns get sql column name.获取数据库列名
var ConfigColumns = struct {
	ConfigID    string
	ConfigName  string
	ConfigKey   string
	ConfigValue string
	ConfigType  string
	CreateBy    string
	CreateTime  string
	UpdateBy    string
	UpdateTime  string
	Remark      string
}{
	ConfigID:    "config_id",
	ConfigName:  "config_name",
	ConfigKey:   "config_key",
	ConfigValue: "config_value",
	ConfigType:  "config_type",
	CreateBy:    "create_by",
	CreateTime:  "create_time",
	UpdateBy:    "update_by",
	UpdateTime:  "update_time",
	Remark:      "remark",
}
