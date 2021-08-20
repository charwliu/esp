package api

import (
	"github.com/gofiber/fiber/v2"
)

func CreateToken(router fiber.Router) {
	router.Post("/", func(ctx *fiber.Ctx) error {
		return Ok(ctx, "ok")
	})
}

func InfoToken(router fiber.Router) {
	router.Get("/", func(ctx *fiber.Ctx) error {
		return Created(ctx, "ok")
	})
}

func DeleteToken(router fiber.Router) {
	router.Delete("/", func(ctx *fiber.Ctx) error {
		return Ok(ctx, "ok")
	})
}
