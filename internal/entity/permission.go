package entity

import (
	"github.com/google/uuid"
)

type Permission struct {
	DefaultModel
	Name          string `gorm:"size:255" json:"name"`
	Description   string `json:"description"`
	IsInternal    bool
	Action        string `gorm:"size:255"`
	Resource      string `gorm:"size:255"`
	Xml           string
	OauthClientId *uuid.UUID
	OauthClient   *OauthClient
}
