package web

import (
	"errors"

	"github.com/google/uuid"

	"go.vixal.xyz/esp/app/models"
	"go.vixal.xyz/esp/platform/database"
)

var zero uuid.UUID

// FindUserByUsername Return a single user as JSON
func FindUserByUsername(db *database.Database, username string) (*models.User, error) {
	User := new(models.User)
	if response := db.Where("name = ?", username).First(&User); response.Error != nil {
		return nil, response.Error
	}
	if User.ID == zero {
		return User, errors.New("user not found")
	}
	// Match role to user
	if User.RoleID != zero {
		Role := new(models.Role)
		if res := db.Find(&Role, "id = ?", User.RoleID); res.Error != nil {
			return User, errors.New("error when retrieving the role of the user")
		}
		if Role.ID != zero {
			User.Role = *Role
		}
	}
	return User, nil
}

// FindUserByID Return a single user as JSON
func FindUserByID(db *database.Database, id int64) (*models.User, error) {
	User := new(models.User)
	if response := db.Where("id = ?", id).First(&User); response.Error != nil {
		return nil, response.Error
	}
	if User.ID == zero {
		return User, errors.New("user not found")
	}
	// Match role to user
	if User.RoleID != zero {
		Role := new(models.Role)
		if res := db.Find(&Role, "id = ?", User.RoleID); res.Error != nil {
			return User, errors.New("error when retrieving the role of the user")
		}
		if Role.ID != zero {
			User.Role = *Role
		}
	}
	return User, nil
}
