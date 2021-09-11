package entity

import (
	"sync"
	"time"

	"github.com/gosimple/slug"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"go.vixal.xyz/esp/internal/event"
	"go.vixal.xyz/esp/pkg/rnd"
	"go.vixal.xyz/esp/pkg/txt"
)

var labelMutex = sync.Mutex{}

type Labels []Label

// Label is used for photo, album and location categorization
type Label struct {
	ID               uint       `gorm:"primary_key" json:"ID" yaml:"-"`
	LabelUID         string     `gorm:"size:42;unique_index;" json:"UID" yaml:"UID"`
	LabelSlug        string     `gorm:"size:255;unique_index;" json:"Slug" yaml:"-"`
	CustomSlug       string     `gorm:"size:255;index;" json:"CustomSlug" yaml:"-"`
	LabelName        string     `gorm:"size:255;" json:"Name" yaml:"Name"`
	LabelPriority    int        `json:"Priority" yaml:"Priority,omitempty"`
	LabelFavorite    bool       `json:"Favorite" yaml:"Favorite,omitempty"`
	LabelDescription string     `gorm:"type:TEXT;" json:"Description" yaml:"Description,omitempty"`
	LabelNotes       string     `gorm:"type:TEXT;" json:"Notes" yaml:"Notes,omitempty"`
	Categories       []*Label   `gorm:"many2many:categories;association_jointable_foreignkey:category_id" json:"-" yaml:"-"`
	CreatedAt        time.Time  `json:"CreatedAt" yaml:"-"`
	UpdatedAt        time.Time  `json:"UpdatedAt" yaml:"-"`
	DeletedAt        *time.Time `sql:"index" json:"DeletedAt,omitempty" yaml:"-"`
	New              bool       `gorm:"-" json:"-" yaml:"-"`
}

func (Label) TableName() string {
	return "label"
}

// BeforeCreate creates a random UID if needed before inserting a new row to the database.
func (m *Label) BeforeCreate(*gorm.DB) error {
	if rnd.IsUID(m.LabelUID, 'l') {
		return nil
	}
	m.LabelUID = rnd.PPID('l')
	return nil
}

// NewLabel returns a new label.
func NewLabel(name string, priority int) *Label {
	labelName := txt.Clip(name, txt.ClipDefault)

	if labelName == "" {
		labelName = "Unknown"
	}

	labelName = txt.Title(labelName)
	labelSlug := slug.Make(txt.Clip(labelName, txt.ClipSlug))

	result := &Label{
		LabelSlug:     labelSlug,
		CustomSlug:    labelSlug,
		LabelName:     labelName,
		LabelPriority: priority,
	}

	return result
}

// Save updates the existing or inserts a new label.
func (m *Label) Save() error {
	labelMutex.Lock()
	defer labelMutex.Unlock()

	return DB().Save(m).Error
}

// Create inserts the label to the database.
func (m *Label) Create() error {
	labelMutex.Lock()
	defer labelMutex.Unlock()

	return DB().Create(m).Error
}

// Delete removes the label from the database.
func (m *Label) Delete() error {
	DB().Where("label_id = ? OR category_id = ?", m.ID, m.ID).Delete(&Category{})
	// DB().Where("label_id = ?", m.ID).Delete(&PhotoLabel{})
	return DB().Delete(m).Error
}

// Deleted returns true if the label is deleted.
func (m *Label) Deleted() bool {
	return m.DeletedAt != nil
}

// Restore the label from the database.
func (m *Label) Restore() error {
	if m.Deleted() {
		return UnscopedDB().Model(m).Update("DeletedAt", nil).Error
	}
	return nil
}

// Update - Updates a label property in the database.
func (m *Label) Update(attr string, value interface{}) error {
	return UnscopedDB().Model(m).UpdateColumn(attr, value).Error
}

// FirstOrCreateLabel returns the existing label, inserts a new label or nil in case of errors.
func FirstOrCreateLabel(m *Label) *Label {
	result := Label{}

	if err := UnscopedDB().Where("label_slug = ? OR custom_slug = ?", m.LabelSlug, m.CustomSlug).First(&result).Error; err == nil {
		return &result
	} else if createErr := m.Create(); createErr == nil {
		if m.LabelPriority >= 0 {
			event.EntitiesCreated("labels", []*Label{m})
			event.Publish("count.labels", event.Data{
				"count": 1,
			})
		}
		return m
	} else {
		L().Error("label: find or create", zap.String("slug", m.LabelSlug), zap.Error(createErr), zap.Error(err))
	}

	return nil
}

// FindLabel returns an existing row if exists.
func FindLabel(s string) *Label {
	labelSlug := slug.Make(txt.Clip(s, txt.ClipSlug))

	result := Label{}

	if err := DB().Where("label_slug = ? OR custom_slug = ?", labelSlug, labelSlug).First(&result).Error; err == nil {
		return &result
	}

	return nil
}

// AfterCreate sets the New column used for database callback
func (m *Label) AfterCreate(*gorm.DB) error {
	m.New = true
	return nil
}

// SetName changes the label name.
func (m *Label) SetName(name string) {
	newName := txt.Clip(name, txt.ClipDefault)

	if newName == "" {
		return
	}

	m.LabelName = txt.Title(newName)
	m.CustomSlug = slug.Make(txt.Clip(name, txt.ClipSlug))
}

// Links returns all share links for this entity.
func (m *Label) Links() Links {
	return FindLinks("", m.LabelUID)
}
