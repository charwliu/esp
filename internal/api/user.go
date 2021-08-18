package api

import (
	"github.com/gofiber/fiber/v2"

	"go.vixal.xyz/esp/internal/entity"
	"go.vixal.xyz/esp/internal/i18n"
	"go.vixal.xyz/esp/internal/service"
)

// ChangePassword
// PUT /api/v1/users/:uid/password
func ChangePassword(router fiber.Router) {
	router.Put("/:uid/password", func(c *fiber.Ctx) error {
		config := service.Config()

		if config.Public() || config.DisableSettings() {
			return i18n.NewResponse(fiber.StatusForbidden, i18n.ErrPublic)
		}

		return nil
	})
}

// GetAllUsers
// GET /v1/users  Send index of users
func GetAllUsers(router fiber.Router) {
	router.Get("/", func(ctx *fiber.Ctx) error {
		var users entity.Users
		if err := entity.DB().Find(&users).Error; err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal error",
				"code":    fiber.StatusInternalServerError,
			})
		}
		if len(users) > 0 {
			return ctx.Status(fiber.StatusOK).JSON(users)
		} else {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "User not found",
				"code":    fiber.StatusNotFound,
			})
		}
	})
}

// CreateUser
// POST /v1/users  Create user
func CreateUser(router fiber.Router) {
	router.Post("/", func(ctx *fiber.Ctx) error {
		user := &entity.User{}
		if err := ctx.BodyParser(user); err != nil {
			return JSON(ctx, fiber.StatusInternalServerError, i18n.ErrUnexpected, err)
		}
		if result, err := entity.FirstOrCreateUser(user); err != nil {
			return JSON(ctx, fiber.StatusInternalServerError, i18n.ErrUnexpected, err)
		} else {
			return ctx.Status(fiber.StatusOK).JSON(result)
		}

	})
}

func GetUserInfo(router fiber.Router) {
	router.Get("/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if user, err := entity.FindUserByUID(id); err != nil {
			return JSON(ctx, fiber.StatusNotFound, i18n.ErrUserNotFound, err)
		} else {
			return ctx.Status(fiber.StatusOK).JSON(user)
		}
	})
}

func EditUser(router fiber.Router) {
	router.Put("/:id", func(ctx *fiber.Ctx) error {
		var editUser entity.User
		if err := ctx.BodyParser(&editUser); err != nil {
			return JSON(ctx, fiber.StatusInternalServerError, i18n.ErrUnexpected, err)
		}
		id := ctx.Params("id")
		if user, err := entity.FindUserByUID(id); err != nil {
			return JSON(ctx, fiber.StatusNotFound, i18n.ErrUserNotFound, err)
		} else {
			user.UserName = editUser.UserName
			user.Email = editUser.Email
			user.Status = editUser.Status
			user.Avatar = editUser.Avatar
			if err = user.Save(); err != nil {
				return JSON(ctx, fiber.StatusInternalServerError, i18n.ErrUnexpected, err)
			}
			return ctx.Status(fiber.StatusOK).JSON(editUser)
		}
	})
}

func DeleteUser(router fiber.Router) {
	router.Delete("/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if err := entity.DB().Delete(&entity.User{}, "user_uid = ?", id).Error; err == nil {
			return ctx.Status(fiber.StatusNoContent).JSON("User " + id + "destroyed")
		} else {
			return JSON(ctx, fiber.StatusNotFound, i18n.ErrUserNotFound, err)
		}
	})
}
