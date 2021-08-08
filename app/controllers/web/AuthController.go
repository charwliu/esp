package web

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	hashing "github.com/thomasvvugt/fiber-hashing"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func IsAuthenticated(store *session.Store, ctx *fiber.Ctx) (authenticated bool) {
	sess, err := store.Get(ctx)
	if err != nil {
		panic(err)
	}
	// Get User ID from session store
	userID, correct := sess.Get("userid").(int64)
	if !correct {
		userID = 0
	}
	auth := false
	if userID > 0 {
		auth = true
	}
	return auth
}

func ShowLoginForm() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		err := ctx.Render("login", fiber.Map{})
		if err != nil {
			if err2 := ctx.Status(500).SendString(err.Error()); err2 != nil {
				zap.L().Error(err2.Error())
			}
		}
		return err
	}
}

func PostLoginForm(hasher hashing.Driver, store *session.Store, db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		username := ctx.FormValue("username")
		// Find user
		user, err := FindUserByUsername(db, username)
		if err != nil {
			zap.L().Error("Error when finding user",
				zap.String("username", username),
				zap.Error(err))
			return err
		}

		// Check if password matches hash
		if hasher != nil {
			password := ctx.FormValue("password")
			match, err := hasher.MatchHash(password, user.Password)
			if err != nil {
				zap.S().Errorf("Error when matching hash for password: %v", err)
			}
			if match {
				sess, err := store.Get(ctx)
				if err != nil {
					panic(err)
				}
				defer func(sess *session.Session) {
					err := sess.Save()
					if err != nil {
						panic(err)
					}
				}(sess)

				// Set the user ID in the session
				sess.Set("userid", user.ID)
				zap.L().Info("User set in store store with", zap.Stringer("ID", user.ID))
				if err := ctx.SendString("You should be logged in successfully!"); err != nil {
					zap.L().Error(err.Error())
				}
			} else {
				if err := ctx.SendString("The entered details do not match our records."); err != nil {
					zap.L().Error(err.Error())
				}
			}
		} else {
			zap.L().Error("Hash provider was not set")
		}
		return nil
	}
}

func PostLogoutForm(sessionLookup string, store *session.Store, db *gorm.DB) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if IsAuthenticated(store, ctx) {
			sess, err := store.Get(ctx)
			if err != nil {
				panic(err)
			}
			sess.Delete("userid")
			if err := sess.Save(); err != nil {
				zap.L().Error(err.Error())
			}
			// Check if cookie needs to be unset
			split := strings.Split(sessionLookup, ":")
			if strings.ToLower(split[0]) == "cookie" {
				// Unset cookie on client-side
				ctx.Set("Set-Cookie", split[1]+"=; expires=Thu, 01 Jan 1970 00:00:00 GMT; path=/; HttpOnly")
				if err := ctx.SendString("You are now logged out."); err != nil {
					zap.L().Error(err.Error())
				}
				return nil
			}
			return nil
		}
		// TODO: Redirect?
		return nil
	}
}
