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
		err := ctx.BodyParser(user)
		if err != nil {
			return err
		}
		if user = entity.FirstOrCreateUser(user); user == nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "Internal error",
				"code":    fiber.StatusInternalServerError,
			})
		}

		user.Password = ""
		return ctx.Status(fiber.StatusOK).JSON(user)

	})
}

func GetUserInfo(router fiber.Router) {
	router.Get("/:id", func(ctx *fiber.Ctx) error {
		var user *entity.User
		id := ctx.Params("id")
		user = entity.FindUserByUID(id)
		if user == nil {
			return ctx.Status(fiber.StatusNotFound).JSON(UserNotFound())
		}
		return ctx.Status(fiber.StatusOK).JSON(user)
	})
}

func EditUser(router fiber.Router) {
	router.Put("/:id", func(ctx *fiber.Ctx) error {
		var editUser entity.User
		err := ctx.BodyParser(&editUser)
		if err != nil {
			return err
		}
		id := ctx.Params("id")
		user := entity.FindUserByUID(id)
		if user == nil {
			return ctx.Status(fiber.StatusNotFound).JSON(UserNotFound())
		}
		user.UserName = editUser.UserName
		user.Email = editUser.Email
		user.Status = editUser.Status
		user.Avatar = editUser.Avatar
		err = user.Save()
		if err != nil {
			return err
		}
		return ctx.Status(fiber.StatusOK).JSON(editUser)
	})
}

func DeleteUser(router fiber.Router)  {
	router.Delete("/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		err := entity.DB().Delete(&entity.User{}, "user_uid = ?", id ).Error
		if err == nil {
			return ctx.Status(fiber.StatusNoContent).JSON("User "+ id + "destroyed")
		}
		return ctx.Status(fiber.StatusNotFound).JSON(UserNotFound())
	})
}
