package api

import (
	"github.com/gofiber/fiber/v2"
	"go.vixal.xyz/esp/internal/i18n"

	"go.vixal.xyz/esp/internal/api"
	"go.vixal.xyz/esp/internal/entity"
	"go.vixal.xyz/esp/internal/event"
)

var log = event.Log.Sugar()

// GetAllUsers -- Send index of users
// GET /v1/users  -- Send index of users
func GetAllUsers(router fiber.Router) {
	router.Get("/", func(ctx *fiber.Ctx) error {
		var users entity.Users
		if err := entity.DB().Find(&users).Error; err != nil {
			return api.JSON(ctx, fiber.StatusInternalServerError, i18n.ErrUnexpected, err)
		}
		if len(users) > 0 {
			//for _, users := range users {
			//	//users.Urls = struct {
			//	//	OrganizationRolesUrl string `json:"organization_roles_url,omitempty"`
			//	//	RolesUrl             string `json:"roles_url,omitempty"`
			//	//}{
			//	//	OrganizationRolesUrl: "v1/users/" + users.UserUID + "/organization_roles",
			//	//	RolesUrl:             "v1/users/" + users.UserUID + "/roles",
			//	//}
			//}
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"users": users,
			})
		} else {
			return api.JSON(ctx, fiber.StatusNotFound, i18n.ErrUserNotFound, nil)
		}
	})
}

func CreateUser(router fiber.Router) {
	router.Post("/", func(ctx *fiber.Ctx) error {
		user := new(entity.User)
		if err := ctx.BodyParser(user); err != nil {
			log.Errorf("An error occurred when parsing the new users: %v", err)
			return api.JSON(ctx, fiber.StatusInternalServerError, i18n.ErrUnexpected, err)
		}
		if result, err := entity.FirstOrCreateUser(user); err != nil {
			return api.JSON(ctx, fiber.StatusInternalServerError, i18n.ErrUnexpected, err)
		} else {
			return ctx.Status(fiber.StatusCreated).JSON(result)
		}
	})
}

// GetUser Return a single users as JSON
func GetUser(router fiber.Router) {
	router.Get("/:user_id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if user, err := entity.FindUserByUID(id); err == nil {
			return ctx.Status(fiber.StatusOK).JSON(user)
		} else {
			return api.JSON(ctx, fiber.StatusNotFound, i18n.ErrUserNotFound, err)
		}
	})
}

// EditUser Edit a single users
func EditUser(router fiber.Router) {
	router.Put("/:user_id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		editUser := new(entity.User)
		if err := ctx.BodyParser(editUser); err != nil {
			log.Error("An error occurred when parsing the edited users: " + err.Error())
		}
		if user, err := entity.FindUserByUID(id); err != nil {
			return api.JSON(ctx, fiber.StatusNotFound, i18n.ErrUserNotFound, err)
		} else {
			user.UserName = editUser.UserName
			// Match role to users
			// Save users
			if err = user.Save(); err != nil {
				return api.JSON(ctx, fiber.StatusInternalServerError, i18n.ErrUnexpected, err)
			}
			return ctx.Status(fiber.StatusOK).JSON(user)
		}

	})
}

// DeleteUser Delete a single users
func DeleteUser(router fiber.Router) {
	router.Delete("/:user_id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		if user, err := entity.FindUserByUID(id); err != nil {
			return api.JSON(ctx, fiber.StatusNotFound, i18n.ErrUserNotFound, err)
		} else {
			if err = user.Delete(); err != nil {
				return api.JSON(ctx, fiber.StatusInternalServerError, i18n.ErrUnexpected, err)
			}
			return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
				"ID":      id,
				"Deleted": true,
			})
		}
	})
}
