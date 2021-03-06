package entity

import (
	"fmt"
	"time"

	"github.com/gosimple/slug"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"go.vixal.xyz/esp/pkg/rnd"
	"go.vixal.xyz/esp/pkg/txt"
)

type Links []Link

// Link represents a sharing link.
type Link struct {
	LinkUID     string    `gorm:"size:42;primary_key;" json:"UID,omitempty" yaml:"UID,omitempty"`
	ShareUID    string    `gorm:"size:42;unique_index:idx_links_uid_token;" json:"Share" yaml:"Share"`
	ShareSlug   string    `gorm:"size:255;index;" json:"Slug" yaml:"Slug,omitempty"`
	LinkToken   string    `gorm:"size:255;unique_index:idx_links_uid_token;" json:"Token" yaml:"Token,omitempty"`
	LinkExpires int       `json:"Expires" yaml:"Expires,omitempty"`
	LinkViews   uint      `json:"Views" yaml:"-"`
	MaxViews    uint      `json:"MaxViews" yaml:"-"`
	HasPassword bool      `json:"HasPassword" yaml:"HasPassword,omitempty"`
	CanComment  bool      `json:"CanComment" yaml:"CanComment,omitempty"`
	CanEdit     bool      `json:"CanEdit" yaml:"CanEdit,omitempty"`
	CreatedAt   time.Time `deepcopier:"skip" json:"CreatedAt" yaml:"CreatedAt"`
	ModifiedAt  time.Time `deepcopier:"skip" yaml:"ModifiedAt"`
}

func (Link) TableName() string {
	return "link"
}

// BeforeCreate creates a random UID if needed before inserting a new row to the database.
func (m *Link) BeforeCreate(*gorm.DB) error {
	if rnd.IsUID(m.LinkUID, 's') {
		return nil
	}
	m.LinkUID = rnd.PPID('s')
	return nil
}

// NewLink creates a sharing link.
func NewLink(shareUID string, canComment, canEdit bool) Link {
	now := Timestamp()

	result := Link{
		LinkUID:    rnd.PPID('s'),
		ShareUID:   shareUID,
		LinkToken:  rnd.Token(10),
		CanComment: canComment,
		CanEdit:    canEdit,
		CreatedAt:  now,
		ModifiedAt: now,
	}

	return result
}

func (m *Link) Redeem() {
	m.LinkViews += 1

	result := DB().Model(m).UpdateColumn("LinkViews", m.LinkViews)

	if result.RowsAffected == 0 {
		L().Warn("link: failed updating share view counter for", zap.String("linkUID", m.LinkUID))
	}
}

func (m *Link) Expired() bool {
	if m.MaxViews > 0 && m.LinkViews >= m.MaxViews {
		return true
	}

	if m.LinkExpires <= 0 {
		return false
	}

	now := Timestamp()
	expires := m.ModifiedAt.Add(Seconds(m.LinkExpires))

	return now.After(expires)
}

func (m *Link) SetSlug(s string) {
	m.ShareSlug = slug.Make(txt.Clip(s, txt.ClipSlug))
}

func (m *Link) SetPassword(password string) error {
	pw := NewPassword(m.LinkUID, password)

	if err := pw.Save(); err != nil {
		return err
	}

	m.HasPassword = true

	return nil
}

func (m *Link) InvalidPassword(password string) bool {
	if !m.HasPassword {
		return false
	}

	pw := FindPassword(m.LinkUID)

	if pw == nil {
		return password != ""
	}

	return pw.InvalidPassword(password)
}

// Save inserts a new row to the database or updates a row if the primary key already exists.
func (m *Link) Save() error {
	if !rnd.IsPPID(m.ShareUID, 0) {
		return fmt.Errorf("link: invalid share uid (%s)", m.ShareUID)
	}

	if m.LinkToken == "" {
		return fmt.Errorf("link: empty share token")
	}

	m.ModifiedAt = Timestamp()

	return DB().Save(m).Error
}

// Delete Deletes the link.
func (m *Link) Delete() error {
	if m.LinkToken == "" {
		return fmt.Errorf("link: empty share token")
	}

	return DB().Delete(m).Error
}

// FindLink returns an entity pointer if exists.
func FindLink(linkUID string) *Link {
	result := Link{}

	if err := DB().Where("link_uid = ?", linkUID).First(&result).Error; err == nil {
		return &result
	} else {
		L().Error("link: not found", zap.Error(err))
	}

	return nil
}

// FindLinks returns a slice of links for a token and share UID (at least one must be provided).
func FindLinks(token, share string) (result Links) {
	if token == "" && share == "" {
		L().Warn("link: share token and uid must not be empty at the same time (find links)")
		return []Link{}
	}

	q := DB()

	if token != "" {
		q = q.Where("link_token = ?", token)
	}

	if share != "" {
		if rnd.IsPPID(share, 'a') {
			q = q.Where("share_uid = ?", share)
		} else {
			q = q.Where("share_slug = ?", share)
		}
	}

	if err := q.Order("modified_at DESC").Find(&result).Error; err != nil {
		L().Error("link: not found", zap.Error(err))
	}

	return result
}

// FindValidLinks returns a slice of non-expired links for a token and share UID (at least one must be provided).
func FindValidLinks(token, share string) (result Links) {
	for _, link := range FindLinks(token, share) {
		if !link.Expired() {
			result = append(result, link)
		}
	}

	return result
}

// String returns an human readable identifier for logging.
func (m *Link) String() string {
	return m.LinkUID
}
