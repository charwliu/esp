package entity

import (
	"github.com/google/uuid"
)

type PepProxy struct {
	ID            string  `gorm:"primary_key;autoIncrement"`
	Password      string `gorm:"size:40"`
	OauthClient   *OauthClient
	OauthClientId *uuid.UUID
}


func (PepProxy) TableName() string {
	return "pep_proxy"
}
