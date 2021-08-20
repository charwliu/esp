package service

import (
	"sync"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/session"

	"go.vixal.xyz/esp/internal/acl"
	"go.vixal.xyz/esp/internal/entity"
)

var onceSession sync.Once

func initSession() {
	services.Store = session.New(Config().SessionConfig())
	services.Store.RegisterType(new(acl.Role))
	services.Store.RegisterType([]*entity.Role{})
}

func Store() *session.Store {
	onceSession.Do(initSession)
	return services.Store
}

func Session(c *fiber.Ctx) (*session.Session, error) {
	return Store().Get(c)
}
