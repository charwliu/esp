package api

import (
	"github.com/gofiber/fiber/v2"
	"go.vixal.xyz/esp/app/models"
	"go.vixal.xyz/esp/internal/api"
	"go.vixal.xyz/esp/internal/entity"
	"go.vixal.xyz/esp/internal/i18n"
)

// GetAllRoles Return all roles as JSON
func GetAllRoles(router fiber.Router) {
	router.Get("/", func(ctx *fiber.Ctx) error {
		var Roles []entity.Role
		if err := entity.DB().Find(&Roles).Error; err != nil {
			log.Errorf("Error occurred while retrieving roles from the database: %v ", err)
			return api.JSON(ctx, fiber.StatusInternalServerError, i18n.ErrUnexpected, err)
		}
		return ctx.Status(fiber.StatusOK).JSON(Roles)
	})
}

// GetRole Return a single role as JSON
func GetRole(router fiber.Router) {
	router.Get("/:id", func(ctx *fiber.Ctx) error {
		Role := new(entity.Role)
		id := ctx.Params("id")
		if err := entity.DB().Find(&Role, "id = ?", id).Error; err != nil {
			log.Errorf("An error occurred when retrieving the role: %v", err)
			return api.JSON(ctx, fiber.StatusInternalServerError, i18n.ErrUnexpected, err)
		}
		if Role.ID == 0 {
			return api.JSON(ctx, fiber.StatusNotFound, i18n.ErrNotFound, nil)
		}
		return ctx.Status(fiber.StatusOK).JSON(Role)
	})

}

// AddRole Add a single role to the database
func AddRole(router fiber.Router) {
	router.Post("/", func(ctx *fiber.Ctx) error {
		Role := new(models.Role)
		if err := ctx.BodyParser(Role); err != nil {
			log.Errorf("An error occurred when parsing the new role: %v", err)
			return api.JSON(ctx, fiber.StatusInternalServerError, i18n.ErrUnexpected, err)
		}
		if err := entity.DB().Create(&Role).Error; err != nil {
			log.Errorf("An error occurred when storing the new role: %v ", err)
			return api.JSON(ctx, fiber.StatusInternalServerError, i18n.ErrUnexpected, err)
		}
		return ctx.Status(fiber.StatusOK).JSON(Role)
	})
}

// EditRole Edit a single role
func EditRole(router fiber.Router) {
	router.Put("/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		EditRole := new(entity.Role)
		Role := new(entity.Role)
		if err := ctx.BodyParser(EditRole); err != nil {
			log.Errorf("An error occurred when parsing the edited role: %v", err)
			return api.JSON(ctx, fiber.StatusInternalServerError, i18n.ErrUnexpected, err)
		}
		if err := entity.DB().Find(&Role, id).Error; err != nil {
			log.Errorf("An error occurred when retrieving the existing role: %v", err)
			return api.JSON(ctx, fiber.StatusInternalServerError, i18n.ErrUnexpected, err)
		}
		// Role does not exist
		if Role == nil {
			return api.JSON(ctx, fiber.StatusNotFound, i18n.ErrNotFound, nil)
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
		if err := entity.DB().Find(&Role).Error; err != nil {
			log.Errorf("An error occurred when finding the role to be deleted: %v", err)
			return api.JSON(ctx, fiber.StatusInternalServerError, i18n.ErrUnexpected, err)
		}
		Role.Delete()
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"ID":      id,
			"Deleted": true,
		})
	})
}
