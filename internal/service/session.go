package service

import (
	"sync"

	"github.com/gofiber/fiber/v2/middleware/session"

	"go.vixal.xyz/esp/internal/acl"
)

var onceSession sync.Once

func initSession() {
	servcies.Session = session.New(Config().SessionConfig())
	var role acl.Role
	servcies.Session.RegisterType(&role)
}

func Session() *session.Store {
	onceSession.Do(initSession)
	return servcies.Session
}
