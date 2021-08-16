package service

import (
	"github.com/gofiber/fiber/v2/middleware/session"

	"go.vixal.xyz/esp/internal/config"
	"go.vixal.xyz/esp/internal/gosh"
	"go.vixal.xyz/esp/internal/query"
)

var conf *config.Config

var servcies struct {
	Query *query.Query
	Session *session.Store
}

func SetConfig(c *config.Config) {
	if c == nil {
		panic("config is nil")
	}

	conf = c

	gosh.SetConfig(c)
}

func Config() *config.Config {
	if conf == nil {
		panic("config is nil")
	}

	return conf
}

