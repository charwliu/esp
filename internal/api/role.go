package api

import (
	"github.com/gofiber/fiber/v2"

	"go.vixal.xyz/esp/app/models"
	"go.vixal.xyz/esp/internal/entity"
	"go.vixal.xyz/esp/internal/i18n"
)

// GetAllRoles Return all roles as JSONError
func GetAllRoles(router fiber.Router) {
	router.Get("/", func(ctx *fiber.Ctx) error {
		var Roles []entity.Role
		if err := entity.DB().Find(&Roles).Error; err != nil {
			log.Errorf("Error occurred while retrieving roles from the database: %v ", err)
			return Unexpected(ctx, err)
		}
		return Ok(ctx, Roles)
	})
}

// GetRole Return a single role as JSONError
func GetRole(router fiber.Router) {
	router.Get("/:id", func(ctx *fiber.Ctx) error {
		Role := new(entity.Role)
		id := ctx.Params("id")
		if err := entity.DB().Find(&Role, "id = ?", id).Error; err != nil {
			log.Errorf("An error occurred when retrieving the role: %v", err)
			return Unexpected(ctx, err)
		}
		if Role.ID == 0 {
			return NotFound(ctx, nil)
		}
		return Ok(ctx, Role)
	})

}

// AddRole Add a single role to the database
func AddRole(router fiber.Router) {
	router.Post("/", func(ctx *fiber.Ctx) error {
		Role := new(models.Role)
		if err := ctx.BodyParser(Role); err != nil {
			log.Errorf("An error occurred when parsing the new role: %v", err)
			return Unexpected(ctx, err)
		}
		if err := entity.DB().Create(&Role).Error; err != nil {
			log.Errorf("An error occurred when storing the new role: %v ", err)
			return Unexpected(ctx, err)
		}
		return Ok(ctx, Role)
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
			return Unexpected(ctx, err)
		}
		if err := entity.DB().Find(&Role, id).Error; err != nil {
			log.Errorf("An error occurred when retrieving the existing role: %v", err)
			return Unexpected(ctx, err)
		}
		// Role does not exist
		if Role == nil {
			return NotFound(ctx, nil)
		}
		Role.RoleName = EditRole.RoleName
		Role.Description = EditRole.Description
		Role.Save()

		return Ok(ctx, Role)

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
			return JSONError(ctx, fiber.StatusInternalServerError, i18n.ErrUnexpected, err)
		}
		Role.Delete()
		return JSON(ctx, fiber.StatusOK,
			struct {
				Id      string
				Deleted bool
			}{id, true}, i18n.MsgOk)
	})
}
