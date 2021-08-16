package entity

import (
	"time"

	"github.com/google/uuid"
)

type OauthAccessToken struct {
	AccessToken   string `gorm:"primary_key;" json:"access_token,omitempty"`
	Expires       time.Time
	Scope         string
	RefreshToken  string
	Valid         bool
	Extra         string
	OauthClient   *OauthClient
	OauthClientId *uuid.UUID
	UserId        *uuid.UUID
	User          *User
	IotId         string
	Iot           *Iot
}

func (OauthAccessToken) TableName() string {
	return "oauth_access_token"
}
