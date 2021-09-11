package entity

import (
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"

	"go.vixal.xyz/esp/internal/event"
	"go.vixal.xyz/esp/internal/maps"
	"go.vixal.xyz/esp/pkg/s2"
	"go.vixal.xyz/esp/pkg/txt"
)

var cellMutex = sync.Mutex{}

// Cell represents a S2 cell with location data.
type Cell struct {
	ID           string    `gorm:"size:42;primary_key;auto_increment:false;" json:"ID" yaml:"ID"`
	CellName     string    `gorm:"size:255;" json:"Name" yaml:"Name,omitempty"`
	CellCategory string    `gorm:"size:64;" json:"Category" yaml:"Category,omitempty"`
	PlaceID      string    `gorm:"size:42;default:'zz'" json:"-" yaml:"PlaceID"`
	Place        *Place    `gorm:"PRELOAD:true" json:"Place" yaml:"-"`
	CreatedAt    time.Time `json:"CreatedAt" yaml:"-"`
	UpdatedAt    time.Time `json:"UpdatedAt" yaml:"-"`
}

func (m Cell) TableName() string {
	return "cell"
}

// UnknownLocation is PhotoPrism's default location.
var UnknownLocation = Cell{
	ID:           "zz",
	Place:        &UnknownPlace,
	PlaceID:      "zz",
	CellName:     "",
	CellCategory: "",
}

// CreateUnknownLocation creates the default location if not exists.
func CreateUnknownLocation() {
	FirstOrCreateCell(&UnknownLocation)
}

// NewCell creates a location using a token extracted from coordinate
func NewCell(lat, lng float32) *Cell {
	result := &Cell{}

	result.ID = s2.PrefixedToken(float64(lat), float64(lng))

	return result
}

// Find retrieves location data from the database or an external api if not known already.
func (m *Cell) Find(api string) error {
	start := time.Now()
	db := DB()

	if err := db.Preload("Place").First(m, "id = ?", m.ID).Error; err == nil {
		L().Debug("location: found cell", zap.String("id", m.ID))
		return nil
	}

	l := &maps.Location{
		ID: s2.NormalizeToken(m.ID),
	}

	if err := l.QueryApi(api); err != nil {
		return err
	}

	if found := FindPlace(l.PrefixedToken(), l.Label()); found != nil {
		m.Place = found
	} else {
		place := &Place{
			ID:            l.PrefixedToken(),
			PlaceLabel:    l.Label(),
			PlaceCity:     l.City(),
			PlaceProvince: l.State(),
			PlaceCountry:  l.CountryCode(),
			PlaceKeywords: l.KeywordString(),
		}

		if createErr := place.Create(); createErr == nil {
			event.Publish("count.places", event.Data{
				"count": 1,
			})

			L().Info("location: added place",
				zap.String("id", place.ID), zap.Duration("duration", time.Since(start)))

			m.Place = place
		} else if found := FindPlace(l.PrefixedToken(), l.Label()); found != nil {
			m.Place = found
		} else {
			L().Error("location: create place",
				zap.String("id", place.ID),
				zap.Error(createErr))
			m.Place = &UnknownPlace
		}
	}
	m.PlaceID = m.Place.ID
	m.CellName = l.Name()
	m.CellCategory = l.Category()

	cellMutex.Lock()
	defer cellMutex.Unlock()

	if createErr := db.Create(m).Error; createErr == nil {
		L().Info("location: added cell",
			zap.String("id", m.ID), zap.Duration("duration", time.Since(start)))
		return nil
	} else if findErr := db.Preload("Place").First(m, "id = ?", m.ID).Error; findErr != nil {
		L().Error("location: create cell", zap.String("id", m.ID), zap.Error(createErr))
		L().Error("location: find cell", zap.String("id", m.ID), zap.Error(findErr))
		return createErr
	} else {
		L().Debug("location: found cell",
			zap.String("id", m.ID), zap.Duration("duration", time.Since(start)))
	}

	return nil

}

func (m *Cell) Create() error {
	return DB().Create(m).Error
}

// FirstOrCreateCell fetches an existing row, inserts a new row or nil in case of errors.
func FirstOrCreateCell(m *Cell) *Cell {
	if m.ID == "" {
		L().Error("location: cell must not be empty")
		return nil
	}

	if m.PlaceID == "" {
		L().Error("location: place must not be empty (find or create cell)", zap.String("id", m.ID))
		return nil
	}

	result := Cell{}

	if findErr := DB().Where("id = ?", m.ID).Preload("Place").First(&result).Error; findErr == nil {
		return &result
	} else if createErr := m.Create(); createErr == nil {
		return m
	} else if err := DB().Where("id = ?", m.ID).Preload("Place").First(&result).Error; err == nil {
		return &result
	} else {
		L().Error("location: %s (find or create cell)", zap.String("id", m.ID), zap.Error(createErr), zap.Error(findErr))
	}

	return nil
}

// Keywords returns search keywords for a location.
func (m *Cell) Keywords() (result []string) {
	if m.Place == nil {
		L().Error("location: place for cell is nil - you might have found a bug",
			zap.String("id", m.ID))
		return result
	}

	result = append(result, txt.Keywords(txt.ReplaceSpaces(m.City(), "-"))...)
	result = append(result, txt.Keywords(txt.ReplaceSpaces(m.Province(), "-"))...)
	result = append(result, txt.Keywords(txt.ReplaceSpaces(m.CountryName(), "-"))...)
	result = append(result, txt.Keywords(m.Category())...)
	result = append(result, txt.Keywords(m.Name())...)
	result = append(result, txt.Words(m.Place.PlaceKeywords)...)

	result = txt.UniqueWords(result)

	return result
}

// Unknown checks if the location has no id
func (m *Cell) Unknown() bool {
	return m.ID == "" || m.ID == UnknownLocation.ID
}

// Name returns name of location
func (m *Cell) Name() string {
	return m.CellName
}

// NoName checks if the location has no name
func (m *Cell) NoName() bool {
	return m.CellName == ""
}

// Category returns the location category
func (m *Cell) Category() string {
	return m.CellCategory
}

// NoCategory checks id the location has no category
func (m *Cell) NoCategory() bool {
	return m.CellCategory == ""
}

// Label returns the location place label
func (m *Cell) Label() string {
	return m.Place.Label()
}

// City returns the location place city
func (m *Cell) City() string {
	return m.Place.City()
}

// LongCity checks if the city name is more than 16 char
func (m *Cell) LongCity() bool {
	return len(m.City()) > 16
}

// NoCity checks if the location has no city
func (m *Cell) NoCity() bool {
	return m.City() == ""
}

// CityContains checks if the location city contains the text string
func (m *Cell) CityContains(text string) bool {
	return strings.Contains(text, m.City())
}

func (m *Cell) Province() string {
	return m.Place.Province()
}

func (m *Cell) NoProvince() bool {
	return m.Place.Province() == ""
}

// State returns the location place state
func (m *Cell) State() string {
	return m.Place.State()
}

// NoState checks if the location place has no state
func (m *Cell) NoState() bool {
	return m.Place.State() == ""
}

// CountryCode returns the location place country code
func (m *Cell) CountryCode() string {
	return m.Place.CountryCode()
}

// CountryName returns the location place country name
func (m *Cell) CountryName() string {
	return m.Place.CountryName()
}
