package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        *uuid.UUID     `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `sql:"index" json:"deleted_at"`
}

func (base *Base) BeforeCreate(scope *gorm.DB) error {
	if base.ID == nil {
		id := uuid.New()
		base.ID = &id
	}
	return nil
}
