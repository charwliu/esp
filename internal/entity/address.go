package entity

import (
	"fmt"
)

type Addresses []Address

// Address represents a postal address.
type Address struct {
	ID              int     `gorm:"primary_key" json:"ID" yaml:"ID"`
	AddressLat      float32 `gorm:"index;" json:"Lat,omitempty" yaml:"Lat,omitempty"`
	AddressLng      float32 `gorm:"index;" json:"Lng,omitempty" yaml:"Lng,omitempty"`
	AddressLine1    string  `gorm:"size:255;" json:"Line1,omitempty" yaml:"Line1,omitempty"`
	AddressLine2    string  `gorm:"size:255;" json:"Line2,omitempty" yaml:"Line2,omitempty"`
	AddressZip      string  `gorm:"size:32;" json:"Zip,omitempty" yaml:"Zip,omitempty"`
	AddressCity     string  `gorm:"size:128;" json:"City,omitempty" yaml:"City,omitempty"`
	AddressProvince string  `gorm:"size:32;" json:"Province,omitempty" yaml:"Province,omitempty"`
	AddressCountry  string  `gorm:"size:2;default:'zz'" json:"Country,omitempty" yaml:"Country,omitempty"`
	AddressNotes    string  `json:"Notes,omitempty" yaml:"Notes,omitempty"`
}

// TableName the database table name.
func (Address) TableName() string {
	return "address"
}

// UnknownAddress Unknown postal address.
var UnknownAddress = Address{
	ID:             1,
	AddressCountry: UnknownCountry.ID,
}

// CreateUnknownAddress creates the default address if not exists.
func CreateUnknownAddress() {
	FirstOrCreateAddress(&UnknownAddress)
}

// FirstOrCreateAddress returns the existing row, inserts a new row or nil in case of errors.
func FirstOrCreateAddress(m *Address) *Address {
	result := Address{}

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

// String returns an identifier that can be used in logs.
func (m *Address) String() string {
	return fmt.Sprintf("%s, %s %s, %s", m.AddressLine1, m.AddressZip, m.AddressCity, m.AddressCountry)
}

// Unknown returns true if the address is unknown.
func (m *Address) Unknown() bool {
	return m.ID < 0
}

// Create inserts a new row to the database.
func (m *Address) Create() error {
	return DB().Create(m).Error
}

// Save the new row to the database.
func (m *Address) Save() error {
	return DB().Save(m).Error
}
