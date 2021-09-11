package entity

import (
	"github.com/gosimple/slug"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"go.vixal.xyz/esp/internal/event"
	"go.vixal.xyz/esp/internal/maps"
)

// altCountryNames defines mapping between different names for the same country
var altCountryNames = map[string]string{
	"United States of America": "USA",
	"United States":            "USA",
	"China":                    "CN",
	"":                         "Unknown",
}

// Country represents a country location
type Country struct {
	ID                 string `gorm:"size:2;primary_key" json:"ID" yaml:"ID"`
	CountrySlug        string `gorm:"size:255;unique_index;" json:"Slug" yaml:"-"`
	CountryName        string `json:"CountryName" yaml:"CountryName,omitempty"`
	CountryDescription string `gorm:"type:TEXT;" json:"Description,omitempty" yaml:"Description,omitempty"`
	CountryNotes       string `gorm:"type:TEXT;" json:"Notes,omitempty" yaml:"Notes,omitempty"`
	New                bool   `gorm:"-" json:"-" yaml:"-"`
}

// UnknownCountry is defined here to use it as a default
var UnknownCountry = Country{
	ID:          "zz",
	CountryName: maps.CountryNames["zz"],
	CountrySlug: "zz",
}

// CreateUnknownCountry is used to initialize the database with the default country
func CreateUnknownCountry() {
	FirstOrCreateCountry(&UnknownCountry)
}

// NewCountry creates a new country, with default country code if not provided
func NewCountry(countryCode string, countryName string) *Country {
	if countryCode == "" {
		return &UnknownCountry
	}

	if altName, ok := altCountryNames[countryName]; ok {
		countryName = altName
	}

	countrySlug := slug.MakeLang(countryName, "en")

	result := &Country{
		ID:          countryCode,
		CountryName: countryName,
		CountrySlug: countrySlug,
	}

	return result
}

func (Country) TableName() string {
	return "country"
}

// Create inserts a new row to the database.
func (m *Country) Create() error {
	return DB().Create(m).Error
}

// FirstOrCreateCountry returns the existing row, inserts a new row or nil in case of errors.
func FirstOrCreateCountry(m *Country) *Country {
	if cacheData, ok := countryCache.Get(m.ID); ok {
		L().Debug("country: cache hit for", zap.String("id", m.ID))

		return cacheData.(*Country)
	}

	result := Country{}

	if findErr := DB().Where("id = ?", m.ID).First(&result).Error; findErr == nil {
		countryCache.SetDefault(m.ID, &result)
		return &result
	} else if createErr := m.Create(); createErr == nil {
		if !m.Unknown() {
			event.EntitiesCreated("countries", []*Country{m})

			event.Publish("count.countries", event.Data{
				"count": 1,
			})
		}
		countryCache.SetDefault(m.ID, m)
		return m
	} else if err := DB().Where("id = ?", m.ID).First(&result).Error; err == nil {
		countryCache.SetDefault(m.ID, &result)
		return &result
	} else {
		L().Error("country: find or create", zap.String("id", m.ID), zap.Error(createErr), zap.Error(findErr))
	}

	return &UnknownCountry
}

// AfterCreate sets the New column used for database callback
func (m *Country) AfterCreate(scope *gorm.DB) error {
	m.New = true
	return nil
}

// Code returns country code
func (m *Country) Code() string {
	return m.ID
}

// Name returns country name
func (m *Country) Name() string {
	return m.CountryName
}

// Unknown returns true if the country is not a known country.
func (m *Country) Unknown() bool {
	return m.ID == "" || m.ID == UnknownCountry.ID
}
