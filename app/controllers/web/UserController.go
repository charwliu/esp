package web

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"go.vixal.xyz/esp/app/models"
)

var zero uuid.UUID

// FindUserByUsername Return a single user as JSON
func FindUserByUsername(db *gorm.DB, username string) (*models.User, error) {
	User := new(models.User)
	if response := db.Where("name = ?", username).First(&User); response.Error != nil {
		return nil, response.Error
	}
	if User.ID == nil {
		return User, errors.New("user not found")
	}
	// Match role to user
	if User.RoleID != nil {
		Role := new(models.Role)
		if res := db.Find(&Role, "id = ?", User.RoleID); res.Error != nil {
			return User, errors.New("error when retrieving the role of the user")
		}
		if Role.ID != nil {
			User.Role = *Role
		}
	}
	return User, nil
}

// FindUserByID Return a single user as JSON
func FindUserByID(db *gorm.DB, id int64) (*models.User, error) {
	User := new(models.User)
	if response := db.Where("id = ?", id).First(&User); response.Error != nil {
		return nil, response.Error
	}
	if User.ID == nil {
		return User, errors.New("user not found")
	}
	// Match role to user
	if User.RoleID != nil {
		Role := new(models.Role)
		if res := db.Find(&Role, "id = ?", User.RoleID); res.Error != nil {
			return User, errors.New("error when retrieving the role of the user")
		}
		if Role.ID != nil {
			User.Role = *Role
		}
	}
	return User, nil
}
