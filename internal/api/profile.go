package api

import (
	"github.com/gofiber/fiber/v2"

	"go.vixal.xyz/esp/internal/entity"
)

// GetProfile
// GET /api/v1/profile
func GetProfile(router fiber.Router)  {
	router.Get("/profile", func(c *fiber.Ctx) error {
		sess := Session(c)
		var user *entity.User
		if sess != nil {
			userid := sess.Get("userid").(string)
			user = entity.FindUserByUID(userid)
		}
		return c.Status(fiber.StatusOK).JSON(user)
	})
}

func GetStats(router fiber.Router) {
	router.Get("/stats", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"avg_processing_time": 10,
		})
	})
}
