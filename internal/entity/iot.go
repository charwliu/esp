package entity

import (
	"github.com/google/uuid"
)

type Iot struct {
	ID            string `gorm:"primary_key;size:32"`
	Password      string `gorm:"size:40"`
	OauthClientId *uuid.UUID
	OauthClient   *OauthClient
}
