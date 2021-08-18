package api

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"go.vixal.xyz/esp/internal/i18n"

	"go.vixal.xyz/esp/internal/api"
	"go.vixal.xyz/esp/internal/entity"
	"go.vixal.xyz/esp/internal/form"
	"go.vixal.xyz/esp/internal/service"
)

// ShowLoginForm method for show Login Form
// @Description Show Login Form
// @Summary show login form
// @Tags auth
// @Produce html
// @Success 200 {string} status "ok"
// @Router /login [get]
func ShowLoginForm(router fiber.Router) {
	router.Get("/login", func(ctx *fiber.Ctx) error {
		err := ctx.Render("login", fiber.Map{})
		if err != nil {
			if err2 := ctx.Status(500).SendString(err.Error()); err2 != nil {
				log.Error(err2.Error())
			}
		}
		return err
	})
}

type Data struct {
	User   *entity.User `json:"users"`  // Store users, guest or anonymous person.
	Tokens []string     `json:"tokens"` // Slice of secret share tokens.
}

// UserLogin
// @Description Show Login Form
// @Summary show login form
// @Tags auth
// @Produce json
// @Accept json
// @Param login body form.Login true "Credentials to use"
// @Success 200 {object} form.User
// @Failure 401 {string} "Unauthorized"
// @Failure 422 {object} i18n.Response
// @Router /api/v1/login [post]
func UserLogin(router fiber.Router) {
	router.Post("/login", func(ctx *fiber.Ctx) error {
		f := form.Login{}
		if err := ctx.BodyParser(&f); err != nil {
			return api.JSON(ctx, fiber.StatusInternalServerError, i18n.ErrUnexpected, err)
		}
		var data Data
		session, err := service.Session(ctx)
		if err != nil {
			return api.JSON(ctx, fiber.StatusInternalServerError, i18n.ErrUnexpected, err)
		}

		//conf := service.Config()

		if f.HasToken() {

		} else if f.HasCredentials() {
			if user, err := entity.FindUserByName(f.UserName); err == nil {
				if !user.InvalidPassword(f.Password) {
					data.User = user
					id := session.ID()
					// Set the users ID in the session
					session.Set("userid", data.User.UserUID)
					//session.Set("role", data.User.Role())
					log.Info("User set in store with", zap.String("ID", data.User.UserUID))
					if err := session.Save(); err != nil {
						return api.JSON(ctx, fiber.StatusInternalServerError, i18n.ErrUnexpected, err)
					}
					role := "users"
					if len(data.User.Roles) > 1 {
						role = data.User.Roles[0].RoleName
					}
					return ctx.Status(fiber.StatusOK).JSON(Response{
						Id:               id,
						Status:           "ok",
						CurrentAuthority: role,
					})
				}
			}
			return api.JSON(ctx, fiber.StatusNotFound, i18n.ErrUserNotFound, err)
		}
		return api.JSON(ctx, fiber.StatusUnauthorized, i18n.ErrInvalidPassword, nil)
	})
}

type Response struct {
	Status           string `json:"status"`
	Type             string `json:"type"`
	CurrentAuthority string `json:"currentAuthority"`
	Id               string `json:"id"`
}

// GetCurrentUser
// GET /api/v1/currentUser
func GetCurrentUser(router fiber.Router) {
	router.Get("/currentUser", func(ctx *fiber.Ctx) error {
		session, err := service.Session(ctx)
		if err != nil {
			return api.JSON(ctx, fiber.StatusUnauthorized, i18n.ErrInvalidCredentials, err)
		}
		if userId, ok := session.Get("userid").(string); ok {
			if user, err := entity.FindUserByUID(userId); err == nil {
				return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
					"success": true,
					"data":    user,
				})
			}
		}
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": true,
			"data": fiber.Map{
				"isLogin": false,
			},
			"errorCode":    "401",
			"errorMessage": "请先登录！",
		})
	})
}

func UserLogout(router fiber.Router) {
	router.Post("/logout", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id", "")
		if id != "" {
			ctx.Request().Header.SetCookie("session_id", id)
		}
		session, err := service.Session(ctx)
		if err != nil {
			return api.JSON(ctx, fiber.StatusUnauthorized, i18n.ErrInvalidCredentials, err)
		}
		id = session.ID()
		if err = session.Destroy(); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": err.Error(),
				"id":     id,
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
			"id":     id,
		})
	})
}

// DeleteSession
// DELETE /api/v1/session/:id
func DeleteSession(router fiber.Router) {
	router.Delete("/session/:id", func(ctx *fiber.Ctx) error {
		if id := ctx.Params("id", ""); id != "" {
			ctx.Request().Header.SetCookie("session_id", id)
		}
		session, err := service.Session(ctx)
		if err != nil {
			return api.JSON(ctx, fiber.StatusUnauthorized, i18n.ErrInvalidCredentials, err)
		}
		id := session.ID()
		if err = session.Destroy(); err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": err.Error(),
				"id":     id,
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "ok",
			"id":     id,
		})
	})

}
