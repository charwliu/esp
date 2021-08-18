package api

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

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
	User   *entity.User `json:"user"`   // Session user, guest or anonymous person.
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
			return err
		}

		var data Data
		sess, err := service.Session().Get(ctx)
		if err != nil {
			panic(err)
		}

		//conf := service.Config()

		if f.HasToken() {

		} else if f.HasCredentials() {
			// Find user
			user := entity.FindUserByName(f.UserName)
			if user == nil {
				return api.InvalidCredentials()
			}

			// Check if password matches hash
			if user.InvalidPassword(f.Password) {
				return api.InvalidCredentials()
			}
			data.User = user
		} else {
			return api.InvalidPassword()
		}
		id := sess.ID()
		// Set the user ID in the session
		sess.Set("userid", data.User.UserUID)
		//sess.Set("role", data.User.Role())
		log.Info("User set in store store with", zap.String("ID", data.User.UserUID))
		if err := sess.Save(); err != nil {
			panic(err)
		}
		role := "user"
		if len(data.User.Roles) > 1 {
			role = data.User.Roles[0].RoleName
		}
		return ctx.Status(fiber.StatusOK).JSON(Response{
			Id:               id,
			Status:           "ok",
			CurrentAuthority: role,
		})
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
		sess, err := service.Session().Get(ctx)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(
				api.InvalidCredentials())
		}
		userId, ok := sess.Get("userid").(string)
		if ok {
			user := entity.FindUserByUID(userId)
			if user != nil {
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
		sess, err := service.Session().Get(ctx)
		if err != nil {
			return api.InvalidCredentials()
		}
		id = sess.ID()
		err = sess.Destroy()
		if err != nil {
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
		id := ctx.Params("id", "")
		if id != "" {
			ctx.Request().Header.SetCookie("session_id", id)
		}
		sess, err := service.Session().Get(ctx)
		if err != nil {
			return api.InvalidCredentials()
		}
		err = sess.Destroy()
		if err != nil {
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
