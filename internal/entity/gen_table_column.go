package entity

import (
	"time"

	"gorm.io/gorm"
)

// GenTableColumn 代码生成业务表字段
type GenTableColumn struct {
	ID            int64          `gorm:"primary_key;not null" json:"columnId"` // 编号
	TableID       int64          `json:"tableId"`                              // 归属表编号
	ColumnName    string         `gorm:"size:200" json:"columnName"`           // 列名称
	ColumnComment string         `gorm:"size:500" json:"columnComment"`        // 列描述
	ColumnType    string         `gorm:"size:100" json:"columnType"`           // 列类型
	JavaType      string         `gorm:"size:500" json:"javaType"`             // JAVA类型
	JavaField     string         `gorm:"size:200" json:"javaField"`            // JAVA字段名
	IsPk          string         `gorm:"size:1" json:"isPk"`                   // 是否主键（1是）
	IsIncrement   string         `gorm:"size:1" json:"isIncrement"`            // 是否自增（1是）
	IsRequired    string         `gorm:"size:1" json:"isRequired"`             // 是否必填（1是）
	IsInsert      string         `gorm:"size:1" json:"isInsert"`               // 是否为插入字段（1是）
	IsEdit        string         `gorm:"size:1" json:"isEdit"`                 // 是否编辑字段（1是）
	IsList        string         `gorm:"size:1" json:"isList"`                 // 是否列表字段（1是）
	IsQuery       string         `gorm:"size:1" json:"isQuery"`                // 是否查询字段（1是）
	QueryType     string         `gorm:"size:200;default:EQ" json:"queryType"` // 查询方式（等于、不等于、大于、小于、范围）
	HTMLType      string         `gorm:"size:200" json:"htmlType"`             // 显示类型（文本框、文本域、下拉框、复选框、 单选框、日期控件）
	DictType      string         `gorm:"size:200;default:''" json:"dictType"`  // 字典类型
	Sort          int            `json:"sort"`                                 // 排序
	CreatedBy     string         `gorm:"size:64;default:''" json:"createBy"`   // 创建者
	UpdatedBy     string         `gorm:"size:64;default:''" json:"updateBy"`   // 更新者
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Remark        string         `gorm:"size:500;default:''" json:"remark"` // 备注信息
	Table         *GenTable
}

// GenTableColumnColumns get sql column name.获取数据库列名
var GenTableColumnColumns = struct {
	ColumnID      string
	TableID       string
	ColumnName    string
	ColumnComment string
	ColumnType    string
	JavaType      string
	JavaField     string
	IsPk          string
	IsIncrement   string
	IsRequired    string
	IsInsert      string
	IsEdit        string
	IsList        string
	IsQuery       string
	QueryType     string
	HTMLType      string
	DictType      string
	Sort          string
	CreateBy      string
	CreateTime    string
	UpdateBy      string
	UpdateTime    string
}{
	ColumnID:      "column_id",
	TableID:       "table_id",
	ColumnName:    "column_name",
	ColumnComment: "column_comment",
	ColumnType:    "column_type",
	JavaType:      "java_type",
	JavaField:     "java_field",
	IsPk:          "is_pk",
	IsIncrement:   "is_increment",
	IsRequired:    "is_required",
	IsInsert:      "is_insert",
	IsEdit:        "is_edit",
	IsList:        "is_list",
	IsQuery:       "is_query",
	QueryType:     "query_type",
	HTMLType:      "html_type",
	DictType:      "dict_type",
	Sort:          "sort",
	CreateBy:      "create_by",
	CreateTime:    "create_time",
	UpdateBy:      "update_by",
	UpdateTime:    "update_time",
}
