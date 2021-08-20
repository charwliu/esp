package api

import (
	"github.com/gofiber/fiber/v2"

	"go.vixal.xyz/esp/internal/acl"
	"go.vixal.xyz/esp/internal/service"
)

// GetConfig
// GET /api/v1/config
func GetConfig(router fiber.Router) {
	router.Get("/config", func(ctx *fiber.Ctx) error {
		sess := Session(ctx)
		if sess == nil {
			return Unauthorized(ctx, nil, nil)
		}
		if err := Auth(sess, acl.ResourceConfig, acl.ActionRead); err != nil {
			return err
		}

		conf := service.Config()
		return Ok(ctx, conf.UserConfig())
	})
}
