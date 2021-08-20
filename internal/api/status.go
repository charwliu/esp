package api

import (
	"github.com/gofiber/fiber/v2"
)

// GetStatus
// GET /api/v1/status
func GetStatus(router fiber.Router) {
	router.Get("/status", func(ctx *fiber.Ctx) error {
		return Ok(ctx, "operational")
	})
}
