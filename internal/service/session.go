package service

import (
	"github.com/gofiber/fiber/v2"
	"sync"

	"github.com/gofiber/fiber/v2/middleware/session"

	"go.vixal.xyz/esp/internal/acl"
)

var onceSession sync.Once

func initSession() {
	services.Store = session.New(Config().SessionConfig())
	var role acl.Role
	services.Store.RegisterType(&role)
}

func Store() *session.Store {
	onceSession.Do(initSession)
	return services.Store
}

func Session(c *fiber.Ctx) (*session.Session, error) {
	return Store().Get(c)
}
