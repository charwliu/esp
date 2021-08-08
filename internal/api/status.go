package api

import (
	"github.com/gofiber/fiber/v2"
)

// GetStatus
// GET /api/v1/status
func GetStatus() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status": "operational",
		})
		return nil
	}
}
