package entity

import (
	"time"

	"gorm.io/gorm"
)

// GenTable 代码生成业务表
type GenTable struct {
	ID             int64          `gorm:"primary_key;not null" json:"tableId"`      // 编号
	TableName      string         `gorm:"size:200;default:''" json:"tableName"`     // 表名称
	TableComment   string         `gorm:"size:500;default:''" json:"tableComment"`  // 表描述
	ClassName      string         `gorm:"size:100;default:''" json:"className"`     // 实体类名称
	TplCategory    string         `gorm:"size:200;default:crud" json:"tplCategory"` // 使用的模板（crud单表操作 tree树表操作）
	PackageName    string         `gorm:"size:100" json:"packageName"`              // 生成包路径
	ModuleName     string         `gorm:"size:30" json:"moduleName"`                // 生成模块名
	BusinessName   string         `gorm:"size:30" json:"businessName"`              // 生成业务名
	FunctionName   string         `gorm:"size:50" json:"functionName"`              // 生成功能名
	FunctionAuthor string         `gorm:"size:50" json:"functionAuthor"`            // 生成功能作者
	Options        string         `gorm:"size:1000" json:"options"`                 // 其它生成选项
	CreatedBy      string         `gorm:"size:64;default:''" json:"createBy"`       // 创建者
	UpdatedBy      string         `gorm:"size:64;default:''" json:"updateBy"`       // 更新者
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Remark         string         `gorm:"size:500;default:''" json:"remark"` // 备注信息
}
