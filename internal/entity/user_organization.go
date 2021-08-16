package entity

import (
	"github.com/google/uuid"
)

type UserOrganization struct {
	ID             int64  `gorm:"primary_key;autoIncrement"`
	Role           string `gorm:"size:10"`
	UserId         *uuid.UUID
	User           *User
	Organization   *Organization
	OrganizationId *uuid.UUID
}

func (UserOrganization) TableName() string {
	return "user_organization"
}
