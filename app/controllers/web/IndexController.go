package web

import (
	"github.com/gofiber/fiber/v2"

	"go.vixal.xyz/esp/internal/event"
)

var log = event.Log.Sugar()

func Index(router fiber.Router) {
	router.Get("/", func(ctx *fiber.Ctx) error {

		//sess := api.Store(ctx)
		//// Bind data to template
		//bind := fiber.Map{
		//	"name": "Fiber",
		//}
		//if sess == nil {
		//	return api.InvalidCredentials()
		//}
		//
		//// Get User ID from sess store
		//userID, ok := sess.Get("userid").(string)
		//if !ok {
		//	return api.InvalidCredentials()
		//}
		//if err := api.Auth(sess, acl.ResourceDefault, acl.ActionRead); err != nil {
		//	return err
		//}
		//user := entity.FindUserByUID(userID)
		//if user == nil {
		//	return api.InvalidCredentials()
		//}
		//bind["username"] = user.UserName
		////bind["auth"] = user.Role()
		//// Render template
		//err := ctx.Render("index")
		//if err != nil {
		//	err2 := ctx.Status(500).SendString(err.Error())
		//	if err2 != nil {
		//		log.Error(err2)
		//	}
		//}
		return nil
	})
}
