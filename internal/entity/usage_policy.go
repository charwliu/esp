package entity

import (
	"time"

	"github.com/google/uuid"
)

type UsagePolicy struct {
	ID            *uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name          string     `gorm:"size:255"`
	Description   string
	Type          string `gorm:"size:20"`
	Parameters    string `gorm:"size:2000"`
	Punishment    string `gorm:"size:20"`
	From          time.Time
	To            time.Time
	Odrl          string
	OauthClientId *uuid.UUID
	OauthClient   OauthClient
}

func (UsagePolicy) TableName() string {
	return "usage_policy"
}
