package entity

import (
	"fmt"
	"time"
)

// Dept 部门表
type Dept struct {
	ID        int64     `gorm:"primary_key" json:"id"`     // 部门id
	Ancestors string    `gorm:"size:50" json:"ancestors"`  // 祖级列表
	DeptName  string    `gorm:"size:32" json:"deptName"`   // 部门名称
	OrderNum  int       `gorm:"default:0" json:"orderNum"` // 显示顺序
	Leader    string    `gorm:"size:32" json:"leader"`     // 负责人
	Phone     string    `gorm:"size:11" json:"phone"`      // 联系电话
	Email     string    `gorm:"size:50" json:"email"`      // 邮箱
	Status    string    `gorm:"size:1" json:"status"`      // 部门状态（0正常 1停用）
	Roles     []Role    `gorm:"many2many:role_dept"`
	CreatedBy string    `gorm:"size:64;default:''" json:"createBy"` // 创建者
	UpdatedBy string    `gorm:"size:64;default:''" json:"updateBy"` // 更新者
	CreatedAt time.Time `json:"created_at"`                         // 创建时间
	UpdatedAt time.Time `json:"updated_at"`                         // 更新时间
	Remark    string    `gorm:"size:500;default:''" json:"remark"`  // 备注信息
	ParentID  int64     `json:"parentId"`                           // 父部门id
	Parent    *Dept
}

// TableName get sql table name.
//  获取数据库表名
func (Dept) TableName() string {
	return "dept"
}

// UnknownDept Unknown postal address.
var UnknownDept = Dept{
	ID:        1,
	Ancestors: "0",
	DeptName:  "老干部服务平台",
	ParentID:  0,
}

// CreateUnknownDept creates the default address if not exists.
func CreateUnknownDept() {
	FirstOrCreateDept(&UnknownDept)
}

func FirstOrCreateDept(m *Dept) *Dept {
	result := Dept{}

	if err := DB().First(&result).Error; err == nil {
		return &result
	} else if createErr := m.Create(); createErr == nil {
		return m
	} else if err := DB().First(&result).Error; err == nil {
		return &result
	} else {
		log.Errorf("address: %s (find or create %s)", createErr, m.String())
	}

	return nil
}

// Unknown returns true if the address is unknown.
func (m *Dept) Unknown() bool {
	return m.ID < 0
}

// Create inserts a new row to the database.
func (m *Dept) Create() error {
	return DB().Create(m).Error
}

// Save the new row to the database.
func (m *Dept) Save() error {
	return DB().Save(m).Error
}

// String returns an identifier that can be used in logs.
func (m *Dept) String() string {
	return fmt.Sprintf("%s:%s", m.Ancestors, m.DeptName)
}
