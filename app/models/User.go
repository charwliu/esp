package models

import (
	"github.com/google/uuid"
)

// User model
type User struct {
	Base
	Name     string `json:"name" xml:"name" form:"name" query:"name"`
	Password string `json:"-" xml:"-" form:"-" query:"-"`
	Email    string
	RoleID   *uuid.UUID `gorm:"type:uuid;column:role_id" json:"role_id"`
	Role     Role       `gorm:"foreignKey:RoleID;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
}
