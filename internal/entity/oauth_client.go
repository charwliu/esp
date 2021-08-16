package entity

import (
	"github.com/google/uuid"
)

type OauthClient struct {
	DefaultModel
	Name         string `gorm:"size:255;not null"`
	Description  string
	Secret       *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"secret"`
	Url          string     `gorm:"size:2000;not null"`
	Image        string     `gorm:"default:'default'"`
	GrantType    string
	ResponseType string
	ClientType   string `gorm:"size:15"`
	Scope        string `gorm:"size:80"`
	Extra        string `gorm:"type:json"`
}

func (OauthClient) TableName() string {
	return "oauth_client"
}
