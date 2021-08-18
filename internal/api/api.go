package api

import (
	"github.com/gofiber/fiber/v2"

	"go.vixal.xyz/esp/internal/event"
	"go.vixal.xyz/esp/internal/i18n"
)

var log = event.Log.Sugar()

func BadRequest() i18n.Response {
	return i18n.NewResponse(fiber.StatusBadRequest, i18n.ErrBadRequest)
}

func EntityNotFound() i18n.Response {
	return i18n.NewResponse(fiber.StatusNotFound, i18n.ErrNotFound)

}

func UserNotFound() i18n.Response {
	return i18n.NewResponse(fiber.StatusNotFound, i18n.ErrUserNotFound)
}

func Unexpected() i18n.Response {
	return i18n.NewResponse(fiber.StatusInternalServerError, i18n.ErrUnexpected)
}

func Abort(code int, id i18n.Message, params ...interface{}) i18n.Response {
	return i18n.NewResponse(code, id, params)
}

func JSON(ctx *fiber.Ctx, code int, id i18n.Message, err error, params ...interface{}) error {
	return ctx.Status(code).JSON(Error(code, id, err, params))
}

func Error(code int, id i18n.Message, err error, params ...interface{}) i18n.Response {
	resp := i18n.NewResponse(code, id, params)
	if err != nil {
		resp.Details = err.Error()
	}
	return resp
}

func Unauthorized() i18n.Response {
	return Abort(fiber.StatusUnauthorized, i18n.ErrUnauthorized)
}

func SaveFailed() i18n.Response {
	return i18n.NewResponse(fiber.StatusInternalServerError, i18n.ErrSaveFailed)
}

func DeleteFailed() i18n.Response {
	return i18n.NewResponse(fiber.StatusInternalServerError, i18n.ErrDeleteFailed)
}

func AlreadyExist(s string) i18n.Response {
	return i18n.NewResponse(fiber.StatusConflict, i18n.ErrAlreadyExists, s)
}

func FeatureDisabled() i18n.Response {
	return i18n.NewResponse(fiber.StatusForbidden, i18n.ErrFeatureDisabled)
}

func InvalidCredentials() i18n.Response {
	return i18n.NewResponse(fiber.StatusBadRequest, i18n.ErrInvalidCredentials)
}

func InvalidPassword() i18n.Response {
	return i18n.NewResponse(fiber.StatusBadRequest, i18n.ErrInvalidPassword)
}
