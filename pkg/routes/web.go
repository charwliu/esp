package routes

import (
	"github.com/gofiber/fiber/v2"
	hashing "github.com/thomasvvugt/fiber-hashing"
	"go.uber.org/zap"

	. "go.vixal.xyz/esp/internal/api"
)

var logger *zap.SugaredLogger

func S() *zap.SugaredLogger {
	if logger == nil {
		logger = zap.S().Named("routes")
	}
	return logger
}

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
			S().Fatalf("Error when creating hash: %v", err)
		}
		if err := ctx.SendString(hash); err != nil {
			panic(err.Error())
		}
		return err
	})

	api := web.Group("/api")
	v1 := api.Group("/v1")
	ShowLoginForm(v1)
	UserLogin(v1)
	UserLogout(v1)
	DeleteSession(v1)
	GetCurrentUser(v1)
}
