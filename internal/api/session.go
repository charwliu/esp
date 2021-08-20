package api

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"go.vixal.xyz/esp/internal/acl"
	"go.vixal.xyz/esp/internal/i18n"
	"go.vixal.xyz/esp/internal/service"
)

// Session returns the current session data.
func Session(c *fiber.Ctx) *session.Session {
	sess, err := service.Session(c)
	if err != nil {
		return nil
	}
	// Check if session id is valid.
	return sess
}

// Auth returns the session if user is authorized for the current action.
func Auth(sess *session.Session, resource acl.Resource, action acl.Action) error {
	role, ok := sess.Get("Role").(acl.Role)
	if ok && acl.Permissions.Deny(resource, role, action) {
		return errors.New(i18n.Msg(i18n.ErrUnauthorized))
	}
	return nil
}
