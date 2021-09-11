package api

import (
	"github.com/gofiber/fiber/v2"

	"go.vixal.xyz/esp/internal/entity"
	"go.vixal.xyz/esp/internal/form"
	"go.vixal.xyz/esp/internal/i18n"
	"go.vixal.xyz/esp/internal/service"
)

// ChangePassword
// PUT /api/v1/users/:uid/password
func ChangePassword(router fiber.Router) {
	router.Put("/:uid/password", func(c *fiber.Ctx) error {
		config := service.Config()

		if config.Public() || config.DisableSettings() {
			return NewResponse(fiber.StatusForbidden, i18n.ErrPublic)
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
			return Unexpected(ctx, err)
		}
		if len(users) > 0 {
			return Success(ctx, users)
		} else {
			return UserNotFound(ctx, nil)
		}
	})
}

// CreateUser
// POST /v1/users  Create user
func CreateUser(router fiber.Router) {
	router.Post("/", func(ctx *fiber.Ctx) error {
		req := &form.User{}
		if err := ctx.BodyParser(req); err != nil {
			return Unexpected(ctx, err)
		}
		if len(req.Username) == 0 {
			req.Username = req.Email
		}
		user := &entity.User{
			UserName:    req.Username,
			PrimaryEmail:       req.Email,
			PhoneNumber: req.Mobile,
		}
		if result, err := entity.FirstOrCreateUser(user); err != nil {
			if result == nil {
				return Unexpected(ctx, err)
			} else {
				return AlreadyExist(ctx, result.UserName)
			}
		} else {
			password := entity.NewPassword(result.UserUID, req.Password)
			if err = password.Save(); err != nil {
				return Unexpected(ctx, err)
			}
			return Success(ctx, result)
		}

	})
}

func GetUserInfo(router fiber.Router) {
	router.Get("/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if user, err := entity.FindUserByUID(id); err != nil {
			return UserNotFound(ctx, err)
		} else {
			return Success(ctx, user)
		}
	})
}

func EditUser(router fiber.Router) {
	router.Put("/:id", func(ctx *fiber.Ctx) error {
		var editUser entity.User
		if err := ctx.BodyParser(&editUser); err != nil {
			return Unexpected(ctx, err)
		}
		id := ctx.Params("id")
		if user, err := entity.FindUserByUID(id); err != nil {
			return UserNotFound(ctx, err)
		} else {
			user.UserName = editUser.UserName
			user.PrimaryEmail = editUser.PrimaryEmail
			user.Status = editUser.Status
			user.Avatar = editUser.Avatar
			if err = user.Save(); err != nil {
				return Unexpected(ctx, err)
			}
			return Success(ctx, editUser)
		}
	})
}

func DeleteUser(router fiber.Router) {
	router.Delete("/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if err := entity.DB().Delete(&entity.User{}, "user_uid = ?", id).Error; err == nil {
			return NoContent(ctx, "User "+id+"destroyed")
		} else {
			return UserNotFound(ctx, err)
		}
	})
}
