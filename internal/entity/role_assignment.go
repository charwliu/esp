package entity

import (
	"github.com/google/uuid"
)

type RoleAssignment struct {
	ID               int64 `gorm:"primary_key;autoIncrement"`
	RoleOrganization string
	OauthClientId    *uuid.UUID
	OauthClient      *OauthClient
	RoleId           string `gorm:"size:36"`
	Role             *Role
	OrganizationId   *uuid.UUID
	Organization     *Organization
	UserId           *uuid.UUID
	User             *User
}
