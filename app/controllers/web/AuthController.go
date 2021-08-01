package web

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
	hashing "github.com/thomasvvugt/fiber-hashing"
	"go.uber.org/zap"

	"go.vixal.xyz/esp/platform/database"
)

func IsAuthenticated(session *session.Session, ctx *fiber.Ctx) (authenticated bool) {
	store := session.Get(ctx)
	// Get User ID from session store
	userID, correct := store.Get("userid").(int64)
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

func PostLoginForm(hasher hashing.Driver, sess *session.Session, db *database.Database) fiber.Handler {
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
				store := sess.Get(ctx)
				defer func(store *session.Store) {
					err := store.Save()
					if err != nil {
						return
					}
				}(store)
				// Set the user ID in the sess store
				store.Set("userid", user.ID)
				zap.L().Info("User set in sess store with", zap.Stringer("ID", user.ID))
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

func PostLogoutForm(sessionLookup string, session *session.Session, db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if IsAuthenticated(session, ctx) {
			store := session.Get(ctx)
			store.Delete("userid")
			if err := store.Save(); err != nil {
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
