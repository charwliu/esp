package api

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"go.vixal.xyz/esp/internal/entity"
	"go.vixal.xyz/esp/internal/form"
	"go.vixal.xyz/esp/internal/i18n"
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
		return ctx.Render("login", fiber.Map{})
	})
}

// UserLogin
// @Description Show Login Form
// @Summary show login form
// @Tags auth
// @Produce json
// @Accept json
// @Param login body form.Login true "Credentials to use"
// @Success 200 {object} form.User
// @Failure 401 {object} api.Response
// @Failure 422 {object} api.Response
// @Failure 500 {object} api.Response
// @Router /api/v1/login [post]
func UserLogin(router fiber.Router) {
	router.Post("/login", func(ctx *fiber.Ctx) error {
		f := form.Login{}
		if err := ctx.BodyParser(&f); err != nil {
			return Unexpected(ctx, err)
		}
		session, err := service.Session(ctx)
		if err != nil {
			return Unexpected(ctx, err)
		}

		// conf := service.Config()

		if f.HasToken() {

		} else if f.HasCredentials() {
			var user *entity.User
			if user, err = entity.FindUserByName(f.UserName); err == nil {
				if !user.InvalidPassword(f.Password) {
					id := session.ID()
					// Set the users ID in the session
					session.Set("userid", user.UserUID)
					session.Set("Role", user.Roles)
					log.Info("User set in store with", zap.String("ID", user.UserUID))
					if err = session.Save(); err != nil {
						return Unexpected(ctx, err)
					}
					role := "users"
					if len(user.Roles) > 1 {
						role = user.Roles[0].RoleName
					}
					return Ok(ctx, struct {
						Id               string `json:"id"`
						CurrentAuthority string `json:"currentAuthority"`
					}{id, role})
				}
			}
			return UserNotFound(ctx, err)
		}
		return InvalidPassword(ctx, nil)
	})
}

// GetCurrentUser
// GET /api/v1/currentUser
func GetCurrentUser(router fiber.Router) {
	router.Get("/currentUser", func(ctx *fiber.Ctx) error {
		session, err := service.Session(ctx)
		if err != nil {
			return JSONError(ctx, fiber.StatusUnauthorized, i18n.ErrInvalidCredentials, err)
		}
		if userId, ok := session.Get("userid").(string); ok {
			if user, err := entity.FindUserByUID(userId); err == nil {
				return Ok(ctx, user)
			}
		}
		return JSON(ctx, fiber.StatusUnauthorized, struct {
			IsLogin bool `json:"isLogin"`
		}{false}, i18n.ErrUnauthorized)
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
			return InvalidCredentials(ctx, err)
		}
		id = session.ID()
		if err = session.Destroy(); err != nil {
			return Unexpected(ctx, err,
				struct {
					Id string `json:"id"`
				}{id})
		}
		return Ok(ctx, struct {
			Id string `json:"id"`
		}{id})
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
			return InvalidCredentials(ctx, err)
		}
		id := session.ID()
		if err = session.Destroy(); err != nil {
			return Unexpected(ctx, err)
		}
		return Ok(ctx, struct {
			Id string `json:"id"`
		}{id})
	})

}
