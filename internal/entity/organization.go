package entity

type Organizations []*Organization

type Organization struct {
	DefaultModel
	Name        string `gorm:"size:64" json:"name"`
	Description string `json:"description"`
	Website     string `gorm:"size:2000" json:"website"`
	Image       string `gorm:"default:'default'" json:"image"`
}


// TableName get sql table name.
//  获取数据库表名
func (Organization) TableName() string {
	return "organization"
}


// Create inserts a new row to the database.
func (m *Organization) Create() error {
	return DB().Create(m).Error
}

// Save the new row to the database.
func (m *Organization) Save() error {
	return DB().Save(m).Error
}

// Delete the row from the database.
func (m *Organization) Delete() error {
	return DB().Delete(m).Error
}
