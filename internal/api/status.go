package api

import (
	"github.com/gofiber/fiber/v2"
)

// GetStatus
// GET /api/v1/status
func GetStatus(router fiber.Router) {
	router.Get("/status", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "operational",
		})
	})
}
