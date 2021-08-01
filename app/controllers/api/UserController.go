package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"go.vixal.xyz/esp/app/models"
	"go.vixal.xyz/esp/pkg/httperror"
	"go.vixal.xyz/esp/platform/database"
)

var zero uuid.UUID

// GetAllUsers Return all users as JSON
func GetAllUsers(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var Users []models.User
		if response := db.Find(&Users); response.Error != nil {
			zap.L().Error("Error occurred while retrieving users from the database: " + response.Error.Error())
			return httperror.Unexpected(response.Error.Error())
		}
		// Match roles to users
		for index, User := range Users {
			if User.RoleID != zero {
				Role := new(models.Role)
				if response := db.Find(&Role, "id = ?", User.RoleID); response.Error != nil {
					zap.L().Error("An error occurred when retrieving the role: " + response.Error.Error())
				}
				if Role.ID != zero {
					Users[index].Role = *Role
				}
			}
		}
		err := ctx.JSON(Users)
		if err != nil {
			zap.L().Error("Error occurred when returning JSON of users: " + err.Error())
		}
		return err
	}
}

// GetUser Return a single user as JSON
func GetUser(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		User := new(models.User)
		id := ctx.Params("id")
		if response := db.Find(&User, "id = ?", id); response.Error != nil {
			zap.L().Error("An error occurred when retrieving the user: " + response.Error.Error())
			return response.Error
		}
		if User.ID == zero {
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				zap.L().Error("Cannot return status not found: " + err.Error())
			}
			err = ctx.JSON(fiber.Map{
				"ID": id,
			})
			if err != nil {
				zap.L().Error("Error occurred when returning JSON of a role: " + err.Error())
			}
			return err
		}
		// Match role to user
		if User.RoleID != zero {
			Role := new(models.Role)
			if response := db.Find(&Role, User.RoleID); response.Error != nil {
				zap.L().Error("An error occurred when retrieving the role: " + response.Error.Error())
			}
			if Role.ID != zero {
				User.Role = *Role
			}
		}
		if err := ctx.JSON(User); err != nil {
			zap.L().Error("Error occurred when returning JSON of a user: " + err.Error())
			return err
		}
		return nil
	}
}

// AddUser Add a single user to the database
func AddUser(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		User := new(models.User)
		if err := ctx.BodyParser(User); err != nil {
			zap.L().Error("An error occurred when parsing the new user: " + err.Error())
			return err
		}
		if response := db.Create(&User); response.Error != nil {
			zap.L().Error("An error occurred when storing the new user: " + response.Error.Error())
			return httperror.Unexpected(response.Error.Error())
		}
		// Match role to user
		if User.RoleID != zero {
			Role := new(models.Role)
			if response := db.Find(&Role, "id = ?", User.RoleID); response.Error != nil {
				zap.L().Error("An error occurred when retrieving the role" + response.Error.Error())
				return response.Error
			}
			if Role.ID != zero {
				User.Role = *Role
			}
		}
		err := ctx.JSON(User)
		if err != nil {
			zap.L().Error("Error occurred when returning JSON of a user: " + err.Error())
		}
		return err
	}
}

// EditUser Edit a single user
func EditUser(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		EditUser := new(models.User)
		User := new(models.User)
		if err := ctx.BodyParser(EditUser); err != nil {
			zap.L().Error("An error occurred when parsing the edited user: " + err.Error())
		}
		if response := db.Find(&User, "id = ?", id); response.Error != nil {
			zap.L().Error("An error occurred when retrieving the existing user: " + response.Error.Error())
		}
		// User does not exist
		if User.ID == zero {
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				zap.L().Error("Cannot return status not found: " + err.Error())
			}
			err = ctx.JSON(fiber.Map{
				"ID": id,
			})
			if err != nil {
				zap.L().Error("Error occurred when returning JSON of a user: " + err.Error())
			}
			return err
		}
		User.Name = EditUser.Name
		User.Email = EditUser.Email
		User.RoleID = EditUser.RoleID
		// Match role to user
		if User.RoleID != zero {
			Role := new(models.Role)
			if response := db.Find(&Role, "id = ?", User.RoleID); response.Error != nil {
				zap.L().Error("An error occurred when retrieving the role" + response.Error.Error())
			}
			if Role.ID != zero {
				User.Role = *Role
			}
		}
		// Save user
		db.Save(&User)

		err := ctx.JSON(User)
		if err != nil {
			zap.L().Error("Error occurred when returning JSON of a user: " + err.Error())
		}
		return err
	}
}

// DeleteUser Delete a single user
func DeleteUser(db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		var User models.User
		db.Find(&User, id)
		if response := db.Find(&User); response.Error != nil {
			zap.L().Error("An error occurred when finding the user to be deleted" + response.Error.Error())
		}
		db.Delete(&User)

		err := ctx.JSON(fiber.Map{
			"ID":      id,
			"Deleted": true,
		})
		if err != nil {
			zap.L().Error("Error occurred when returning JSON of a user: " + err.Error())
		}
		return err
	}
}
