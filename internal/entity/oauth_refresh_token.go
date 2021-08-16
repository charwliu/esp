package entity

import (
	"time"

	"github.com/google/uuid"
)

type OauthRefreshToken struct {
	RefreshToken string `gorm:"primary_key;size:256" json:"access_token,omitempty"`
	Expires      time.Time
	Scope         string
	OauthClient   *OauthClient
	OauthClientId *uuid.UUID
	UserId        *uuid.UUID
	User          *User
	IotId         string
	Iot           *Iot
}

func (OauthRefreshToken) TableName() string {
	return "oauth_refresh_token"
}
