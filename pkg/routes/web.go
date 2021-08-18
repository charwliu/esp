package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	hashing "github.com/thomasvvugt/fiber-hashing"

	. "go.vixal.xyz/esp/app/controllers/api"
)

func RegisterWeb(web fiber.Router, hasher hashing.Driver) {
	// Homepage
	// GET /
	//Controller.Index(web)

	// Panic test route, this brings up an error
	web.Get("/panic", func(ctx *fiber.Ctx) error {
		panic("Hi, I'm a panic error!")
	})

	// Test to load static, compiled assets
	web.Get("/test", func(c *fiber.Ctx) error {
		return c.Render("test", fiber.Map{})
	})

	// Make a new hash
	web.Get("/hash/*", func(ctx *fiber.Ctx) error {
		hash, err := hasher.CreateHash(ctx.Params("*"))
		if err != nil {
			log.Fatalf("Error when creating hash: %v", err)
		}
		if err := ctx.SendString(hash); err != nil {
			panic(err.Error())
		}
		return err
	})

	api := web.Group("/api")
	apiv1 := api.Group("/v1")
	ShowLoginForm(apiv1)
	UserLogin(apiv1)
	UserLogout(apiv1)
	DeleteSession(apiv1)
	GetCurrentUser(apiv1)
}
