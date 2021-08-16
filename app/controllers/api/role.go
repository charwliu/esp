package api

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"go.vixal.xyz/esp/app/models"
	"go.vixal.xyz/esp/internal/api"
	"go.vixal.xyz/esp/internal/entity"
)

// GetAllRoles Return all roles as JSON
func GetAllRoles(router fiber.Router) {
	router.Get("/", func(ctx *fiber.Ctx) error {
		var Roles []entity.Role
		if response := entity.DB().Find(&Roles); response.Error != nil {
			log.Error("Error occurred while retrieving roles from the database: " + response.Error.Error())
			return ctx.Status(fiber.StatusInternalServerError).JSON(api.Unexpected())
		}
		return ctx.Status(fiber.StatusOK).JSON(Roles)
	})
}

// GetRole Return a single role as JSON
func GetRole(router fiber.Router) {
	router.Get("/:id", func(ctx *fiber.Ctx) error {
		Role := new(models.Role)
		id := ctx.Params("id")
		if response := entity.DB().Find(&Role, "id = ?", id); response.Error != nil {
			zap.L().Error("An error occurred when retrieving the role: " + response.Error.Error())
			return ctx.Status(fiber.StatusInternalServerError).JSON(api.Unexpected())
		}
		if Role.ID == nil {
			// Send status not found
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				zap.L().Error("Cannot return status not found: " + err.Error())
				return ctx.Status(fiber.StatusInternalServerError).JSON(api.Unexpected())
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
	})

}

// AddRole Add a single role to the database
func AddRole(router fiber.Router) {
	router.Post("/", func(ctx *fiber.Ctx) error {
		Role := new(models.Role)
		if err := ctx.BodyParser(Role); err != nil {
			zap.L().Error("An error occurred when parsing the new role: " + err.Error())
			return err
		}
		if response := entity.DB().Create(&Role); response.Error != nil {
			zap.L().Error("An error occurred when storing the new role: " + response.Error.Error())
			return response.Error
		}
		err := ctx.JSON(Role)
		if err != nil {
			zap.L().Error("Error occurred when returning JSON of a role: " + err.Error())
		}
		return err
	})
}

// EditRole Edit a single role
func EditRole(router fiber.Router) {
	router.Put("/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		EditRole := new(entity.Role)
		Role := new(entity.Role)
		if err := ctx.BodyParser(EditRole); err != nil {
			zap.L().Error("An error occurred when parsing the edited role: " + err.Error())
		}
		if response := entity.DB().Find(&Role, id); response.Error != nil {
			zap.L().Error("An error occurred when retrieving the existing role: " + response.Error.Error())
			return response.Error
		}
		// Role does not exist
		if Role == nil {
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
		Role.RoleName = EditRole.RoleName
		Role.Description = EditRole.Description
		Role.Save()

		return ctx.Status(fiber.StatusOK).JSON(Role)

	})
}

// DeleteRole Delete a single role
func DeleteRole(router fiber.Router) {
	router.Delete("/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		var Role entity.Role
		entity.DB().Find(&Role, id)
		if response := entity.DB().Find(&Role); response.Error != nil {
			zap.L().Error("An error occurred when finding the role to be deleted: " + response.Error.Error())

		}
	    Role.Delete()

		err := ctx.JSON(fiber.Map{
			"ID":      id,
			"Deleted": true,
		})
		if err != nil {
			zap.L().Error("Error occurred when returning JSON of a role: " + err.Error())
		}
		return err
	})
}
