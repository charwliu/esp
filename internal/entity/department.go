package entity

import (
	"strings"
	"time"

	"go.uber.org/zap"
)

// Department 部门表
type Department struct {
	ID             int64       `gorm:"primary_key" json:"id"`           // 部门id
	DepartmentName string      `gorm:"size:32" json:"DepartmentName"`   // 部门名称
	OrderNum       int         `gorm:"default:0" json:"OrderNum"`       // 显示顺序
	Leader         string      `gorm:"size:32" json:"Leader,omitempty"` // 负责人
	Phone          string      `gorm:"size:11" json:"Phone,omitempty"`  // 联系电话
	Email          string      `gorm:"size:50" json:"Email,omitempty"`  // 邮箱
	Status         string      `gorm:"size:1" json:"Status,omitempty"`  // 部门状态（0正常 1停用）
	Roles          []Role      `gorm:"many2many:role_dept" json:"Roles,omitempty"`
	CreatedBy      string      `gorm:"size:64;default:''" json:"-"`                 // 创建者
	UpdatedBy      string      `gorm:"size:64;default:''" json:"-"`                 // 更新者
	CreatedAt      time.Time   `json:"-"`                                           // 创建时间
	UpdatedAt      time.Time   `json:"-"`                                           // 更新时间
	Remark         string      `gorm:"size:500;default:''" json:"Remark,omitempty"` // 备注信息
	ParentID       int64       `json:"-"`                                           // 父部门id
	Parent         *Department `json:"parent,omitempty"`
	Ancestors      []string    `gorm:"-" json:"Ancestors"`
}

// TableName get sql table name.
//  获取数据库表名
func (Department) TableName() string {
	return "department"
}

// UnknownDepartment Unknown Department.
var UnknownDepartment = Department{
	ID:             1,
	DepartmentName: "老干部服务平台",
	ParentID:       0,
}

// CreateUnknownDepartment creates the default address if not exists.
func CreateUnknownDepartment() {
	FirstOrCreateDepartment(&UnknownDepartment)
}

func FirstOrCreateDepartment(m *Department) *Department {
	result := Department{}

	if err := DB().Where("id = ?", m.ID).First(&result).Error; err == nil {
		return &result
	} else if createErr := m.Create(); createErr == nil {
		return m
	} else {
		L().Error("department: find or create", zap.Stringer("id", m), zap.Error(createErr), zap.Error(err))
	}

	return nil
}

// Unknown returns true if the Department is unknown.
func (m *Department) Unknown() bool {
	return m.ID <= 1
}

// Create inserts a new row to the database.
func (m *Department) Create() error {
	return DB().Create(m).Error
}

// Save the new row to the database.
func (m *Department) Save() error {
	return DB().Save(m).Error
}

func FindDepartmentById(id int64) *Department {
	partial := struct {
		ID             int64
		ParentID       int64
		DepartmentName string
		Ancestors      string
	}{}
	err := DB().Raw(`WITH RECURSIVE cte (id, parent_id, department_name, ancestors) AS (
SELECT id, parent_id, department_name, ARRAY[]::VARCHAR[]
FROM department
WHERE parent_id = 1
UNION ALL
SELECT t.id, t.parent_id, t.department_name, (c.ancestors || t.department_name::VARCHAR)::VARCHAR[]
FROM department t INNER JOIN cte c on t.parent_id = c.id
)
SELECT * FROM cte WHERE parent_id > 1 AND id = ? ORDER BY id;`, id).Scan(&partial).Error
	if err != nil {

	}

	return &Department{
		ID:             partial.ID,
		DepartmentName: partial.DepartmentName,
		ParentID:       partial.ParentID,
		Ancestors:      strings.Split(strings.Trim(partial.Ancestors, "{}"), ","),
	}
}

func (m *Department) String() string {
	if len(m.Ancestors) > 0 {
		return strings.Join(m.Ancestors, "-")
	}
	return m.DepartmentName
}
