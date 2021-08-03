package api

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"go.vixal.xyz/esp/app/models"
	"go.vixal.xyz/esp/pkg/httperror"
)

// GetAllRoles Return all roles as JSON
func GetAllRoles(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var Role []models.Role
		if response := db.Find(&Role); response.Error != nil {
			zap.L().Error("Error occurred while retrieving roles from the database: " + response.Error.Error())
			return ctx.Status(fiber.StatusInternalServerError).JSON(httperror.Unexpected(response.Error.Error()))
		}
		err := ctx.JSON(Role)
		if err != nil {
			zap.L().Error("Error occurred when returning JSON of roles: " + err.Error())
		}
		return err
	}
}

// GetRole Return a single role as JSON
func GetRole(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		Role := new(models.Role)
		id := ctx.Params("id")
		if response := db.Find(&Role, "id = ?", id); response.Error != nil {
			zap.L().Error("An error occurred when retrieving the role: " + response.Error.Error())
			return ctx.Status(fiber.StatusInternalServerError).JSON(httperror.Unexpected(response.Error.Error()))
		}
		if Role.ID == nil {
			// Send status not found
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				zap.L().Error("Cannot return status not found: " + err.Error())
				return ctx.Status(fiber.StatusInternalServerError).JSON(httperror.Unexpected(err.Error()))
			}
			// Set ID
			err = ctx.JSON(fiber.Map{
				"ID": id,
			})
			if err != nil {
				zap.L().Error("Error occurred when returning JSON of a role: " + err.Error())
			}
			return err
		}
		err := ctx.JSON(Role)
		if err != nil {
			zap.L().Error("Error occurred when returning JSON of a role: " + err.Error())
		}
		return err
	}
}

// AddRole Add a single role to the database
func AddRole(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		Role := new(models.Role)
		if err := ctx.BodyParser(Role); err != nil {
			zap.L().Error("An error occurred when parsing the new role: " + err.Error())
			return err
		}
		if response := db.Create(&Role); response.Error != nil {
			zap.L().Error("An error occurred when storing the new role: " + response.Error.Error())
			return response.Error
		}
		err := ctx.JSON(Role)
		if err != nil {
			zap.L().Error("Error occurred when returning JSON of a role: " + err.Error())
		}
		return err
	}
}

// EditRole Edit a single role
func EditRole(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		EditRole := new(models.Role)
		Role := new(models.Role)
		if err := ctx.BodyParser(EditRole); err != nil {
			zap.L().Error("An error occurred when parsing the edited role: " + err.Error())
		}
		if response := db.Find(&Role, id); response.Error != nil {
			zap.L().Error("An error occurred when retrieving the existing role: " + response.Error.Error())
			return response.Error
		}
		// Role does not exist
		if Role.ID == nil {
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				zap.L().Error("Cannot return status not found: " + err.Error())
				return err
			}
			err = ctx.JSON(fiber.Map{
				"ID": id,
			})
			if err != nil {
				zap.L().Error("Error occurred when returning JSON of a role: " + err.Error())
			}
			return err
		}
		Role.Name = EditRole.Name
		Role.Description = EditRole.Description
		db.Save(&Role)

		err := ctx.JSON(Role)
		if err != nil {
			zap.L().Error("Error occurred when returning JSON of a role: " + err.Error())
		}
		return err
	}
}

// DeleteRole Delete a single role
func DeleteRole(db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		var Role models.Role
		db.Find(&Role, id)
		if response := db.Find(&Role); response.Error != nil {
			zap.L().Error("An error occurred when finding the role to be deleted: " + response.Error.Error())

		}
		db.Delete(&Role)

		err := ctx.JSON(fiber.Map{
			"ID":      id,
			"Deleted": true,
		})
		if err != nil {
			zap.L().Error("Error occurred when returning JSON of a role: " + err.Error())
		}
		return err
	}
}
