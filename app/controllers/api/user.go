package api

import (
	"github.com/gofiber/fiber/v2"

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
		err := entity.DB().Find(&users).Error
		if err != nil {
			return api.Unexpected()
		}
		if len(users) > 0 {
			//for _, user := range users {
			//	//user.Urls = struct {
			//	//	OrganizationRolesUrl string `json:"organization_roles_url,omitempty"`
			//	//	RolesUrl             string `json:"roles_url,omitempty"`
			//	//}{
			//	//	OrganizationRolesUrl: "v1/users/" + user.UserUID + "/organization_roles",
			//	//	RolesUrl:             "v1/users/" + user.UserUID + "/roles",
			//	//}
			//}
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"users": users,
			})
		} else {
			return ctx.Status(fiber.StatusNotFound).JSON(api.UserNotFound())
		}
	})
}

func CreateUser(router fiber.Router) {
	router.Post("/", func(ctx *fiber.Ctx) error {
		user := new(entity.User)
		if err := ctx.BodyParser(user); err != nil {
			log.Error("An error occurred when parsing the new user: " + err.Error())
			return err
		}
		if result := entity.FirstOrCreateUser(user); result == nil {
			return api.Unexpected()
		}
		err := ctx.Status(fiber.StatusCreated).JSON(user)
		if err != nil {
			log.Error("Error occurred when returning JSON of a user: " + err.Error())
		}
		return err
	})
}

// GetUser Return a single user as JSON
func GetUser(router fiber.Router) {
	router.Get("/:user_id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		user := entity.FindUserByUID(id)
		if user == nil {
			ctx.Status(fiber.StatusNotFound)
		}
		return ctx.Status(fiber.StatusOK).JSON(user)
	})
}

// EditUser Edit a single user
func EditUser(router fiber.Router) {
	router.Put("/:user_id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		editUser := new(entity.User)
		if err := ctx.BodyParser(editUser); err != nil {
			log.Error("An error occurred when parsing the edited user: " + err.Error())
		}
		user := entity.FindUserByUID(id)
		if user == nil {
			err := ctx.SendStatus(fiber.StatusNotFound)
			if err != nil {
				log.Error("Cannot return status not found: " + err.Error())
			}
			err = ctx.JSON(fiber.Map{
				"ID": id,
			})
			if err != nil {
				log.Error("Error occurred when returning JSON of a user: " + err.Error())
			}
			return err
		}
		user.UserName = editUser.UserName
		// Match role to user
		// Save user
		err := user.Save()
		if err != nil {
			return err
		}
		return ctx.Status(fiber.StatusOK).JSON(user)

	})
}

// DeleteUser Delete a single user
func DeleteUser(router fiber.Router) {
	router.Delete("/:user_id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		user := entity.FindUserByUID(id)
		if user == nil {
			return api.UserNotFound()
		}

		if err := user.Delete(); err != nil {
			return err
		}

		return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"ID":      id,
			"Deleted": true,
		})

	})
}
