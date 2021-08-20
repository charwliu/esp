package api

import (
	"github.com/gofiber/fiber/v2"

	"go.vixal.xyz/esp/internal/i18n"

	"go.vixal.xyz/esp/internal/entity"
)

// GetProfile
// GET /api/v1/profile
func GetProfile(router fiber.Router) {
	router.Get("/profile", func(c *fiber.Ctx) error {
		sess := Session(c)
		if sess != nil {
			if userid, ok := sess.Get("userid").(string); ok {
				if user, err := entity.FindUserByUID(userid); err != nil {
					return c.Status(fiber.StatusOK).JSON(user)
				}
			}
		}
		return JSONError(c, fiber.StatusUnauthorized, i18n.ErrUnauthorized, nil)
	})
}

func GetStats(router fiber.Router) {
	router.Get("/stats", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"avg_processing_time": 10,
		})
	})
}
