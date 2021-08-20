package routes

import (
	"github.com/gofiber/fiber/v2"

	"go.vixal.xyz/esp/internal/api"
)

func RegisterAPI(router fiber.Router) {
	registerRoles(router)
	registerUsers(router)
	registerTokens(router)
	registerGeographic(router)
	api.GetStatus(router)
	api.GetProfile(router)
	api.GetStats(router)
	api.GetNotices(router)
	api.GetFakeList(router)
}

func registerGeographic(router fiber.Router) {
	geo := router.Group("/geographic")
	api.GetProvince(geo)
	api.GetCity(geo)
}

func registerRoles(router fiber.Router) {
	roles := router.Group("/roles")

	api.GetAllRoles(router)

	api.GetRole(roles)
	api.AddRole(roles)
	api.EditRole(roles)
	api.DeleteRole(roles)
}

func registerUsers(router fiber.Router) {
	users := router.Group("/users")
	api.GetAllUsers(users)
	api.GetUserInfo(users)
	api.CreateUser(users)
	api.DeleteUser(users)
	api.EditUser(users)
	api.ChangePassword(users)
}

func registerTokens(router fiber.Router) {
	tokens := router.Group("/tokens")
	api.CreateToken(tokens)
	api.InfoToken(tokens)
	api.DeleteToken(tokens)
}
