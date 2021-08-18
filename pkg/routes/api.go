package routes

import (
	"github.com/gofiber/fiber/v2"

	Controller "go.vixal.xyz/esp/app/controllers/api"
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
	Controller.GetNotices(router)
	Controller.GetFakeList(router)
}

func registerGeographic(router fiber.Router) {
	geo := router.Group("/geographic")
	Controller.GetProvince(geo)
	Controller.GetCity(geo)
}

func registerRoles(router fiber.Router) {
	roles := router.Group("/roles")

	Controller.GetAllRoles(router)

	Controller.GetRole(roles)
	Controller.AddRole(roles)
	Controller.EditRole(roles)
	Controller.DeleteRole(roles)
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
