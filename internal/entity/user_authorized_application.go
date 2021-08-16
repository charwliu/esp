package entity

import (
	"github.com/google/uuid"
)

type UserAuthorizedApplication struct {
	ID            int64      `gorm:"primary_key;autoIncrement"`
	UserId        *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	OauthClientId *uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	OauthClient   *OauthClient
}

func (a UserAuthorizedApplication) TableName() string {
	return "user_authorized_application"
}
