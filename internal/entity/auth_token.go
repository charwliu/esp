package entity

import (
	"time"

	"github.com/google/uuid"
)

type AuthToken struct {
	AccessToken string `gorm:"primary_key;not null;unique"`
	Expires     time.Time
	Valid       bool
	UserId      *uuid.UUID
	User        *User
	PepProxyId  string
	PepProxy    *PepProxy
}

func (AuthToken) TableName() string  {
	return "auth_token"
}
