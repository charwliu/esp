package web

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/session/v2"
	"go.uber.org/zap"

	"go.vixal.xyz/esp/platform/database"
)

func Index(sess *session.Session, db *database.Database) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		auth := IsAuthenticated(sess, ctx)

		// Bind data to template
		bind := fiber.Map{
			"name": "Fiber",
			"auth": auth,
		}

		if auth {
			store := sess.Get(ctx)
			// Get User ID from sess store
			userID, _ := store.Get("userid").(int64)
			user, err := FindUserByID(db, userID)
			if err != nil {
				log.Fatalf("Error when finding user by ID: %v", err)
			}
			bind["username"] = user.Name
		}

		// Render template
		err := ctx.Render("index", bind)
		if err != nil {
			err2 := ctx.Status(500).SendString(err.Error())
			if err2 != nil {
				zap.L().Error(err2.Error())
			}
		}
		return err
	}
}
