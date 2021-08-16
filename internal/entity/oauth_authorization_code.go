package entity

import (
	"time"

	"github.com/google/uuid"
)

type OauthAuthorizationCode struct {
	AuthorizationCode string `gorm:"primary_key;size:256"`
	Expires           time.Time
	RedirectUri       string `gorm:"size:2000"`
	Scope             string
	Valid             bool
	Extra             string
	OauthClient       *OauthClient
	OauthClientId     *uuid.UUID
	UserId            *uuid.UUID
	User              *User
}

func (o OauthAuthorizationCode) TableName() string {
	return "oauth_authorization_code"
}
