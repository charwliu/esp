package api

import (
	"github.com/gofiber/fiber/v2"
)

func CreateToken(router fiber.Router)  {
	router.Post("/", func(ctx *fiber.Ctx) error {
		return nil
	})
}

func InfoToken(router fiber.Router)  {
	router.Get("/", func(ctx *fiber.Ctx) error {
		ctx.Status(fiber.StatusCreated).JSON("ok")
		return nil
	})
}

func DeleteToken(router fiber.Router)  {
	router.Delete("/", func(ctx *fiber.Ctx) error {
		return nil
	})
}
