package entity

import (
	"github.com/google/uuid"
)

type RolePermission struct {
	ID               int64 `gorm:"primary_key;autoIncrement"`
	RoleOrganization string
	PermissionId     *uuid.UUID
	Permission       *Permission
	RoleId           string `gorm:"size:36"`
	Role             *Role
}
