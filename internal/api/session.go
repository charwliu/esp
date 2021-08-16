package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"

	"go.vixal.xyz/esp/internal/acl"
	"go.vixal.xyz/esp/internal/service"

)

// Session returns the current session data.
func Session(c *fiber.Ctx) *session.Session {
	sess, err := service.Session().Get(c)
	if err != nil {
		return nil
	}
	// Check if session id is valid.
	return sess
}


// Auth returns the session if user is authorized for the current action.
func Auth(sess *session.Session, resource acl.Resource, action acl.Action) error {
	role := sess.Get("Role").(acl.Role)
	if acl.Permissions.Deny(resource, role, action) {
		return Unauthorized()
	}
	return nil
}

