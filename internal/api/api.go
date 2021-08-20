package api

import (
	"github.com/gofiber/fiber/v2"

	"go.vixal.xyz/esp/internal/event"
	"go.vixal.xyz/esp/internal/i18n"
)

var log = event.Log.Sugar()

func Success(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusOK).JSON(data)
}

func Ok(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusOK).
		JSON(NewResponseWithData(fiber.StatusOK, data, i18n.MsgOk))
}

func Created(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusCreated).
		JSON(NewResponseWithData(fiber.StatusCreated, data, i18n.MsgOk))
}

func NoContent(ctx *fiber.Ctx, data interface{}) error {
	return ctx.Status(fiber.StatusNoContent).JSON(data)
}

func BadRequest(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusBadRequest).
		JSON(NewResponse(fiber.StatusBadRequest, i18n.ErrBadRequest))
}

func NotFound(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusNotFound).
		JSON(Error(fiber.StatusNotFound, i18n.ErrNotFound, err))
}

func UserNotFound(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusNotFound).
		JSON(Error(fiber.StatusNotFound, i18n.ErrUserNotFound, err))
}

func Unexpected(ctx *fiber.Ctx, err error, params ...interface{}) error {
	var data interface{}
	if len(params) > 0 {
		data = params[0]
	}
	return ctx.Status(fiber.StatusInternalServerError).
		JSON(NewResponseWithDataOrError(fiber.StatusInternalServerError, data, err, i18n.ErrUnexpected))
}

func JSONError(ctx *fiber.Ctx, code int, id i18n.Message, err error, params ...interface{}) error {
	return ctx.Status(code).JSON(Error(code, id, err, params))
}

func JSON(ctx *fiber.Ctx, code int, data interface{}, id i18n.Message, params ...interface{}) error {
	return ctx.Status(code).JSON(NewResponseWithData(code, data, id, params...))
}

func Error(code int, id i18n.Message, err error, params ...interface{}) Response {
	resp := NewResponse(code, id, params...)
	if err != nil {
		resp.Details = err.Error()
	}
	return resp
}

func Unauthorized(ctx *fiber.Ctx, data interface{}, err error) error {
	return ctx.Status(fiber.StatusUnauthorized).
		JSON(NewResponseWithDataOrError(fiber.StatusUnauthorized, data, err, i18n.ErrUnauthorized))
}

func SaveFailed(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).
		JSON(NewResponse(fiber.StatusInternalServerError, i18n.ErrSaveFailed))
}

func DeleteFailed(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusInternalServerError).
		JSON(NewResponse(fiber.StatusInternalServerError, i18n.ErrDeleteFailed))
}

func AlreadyExist(ctx *fiber.Ctx, s string) error {
	return ctx.Status(fiber.StatusConflict).
		JSON(NewResponse(fiber.StatusConflict, i18n.ErrAlreadyExists, s))
}

func FeatureDisabled(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusForbidden).
		JSON(NewResponse(fiber.StatusForbidden, i18n.ErrFeatureDisabled))
}

func InvalidCredentials(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusBadRequest).
		JSON(Error(fiber.StatusBadRequest, i18n.ErrInvalidCredentials, err))
}

func InvalidPassword(ctx *fiber.Ctx, err error) error {
	return ctx.Status(fiber.StatusBadRequest).
		JSON(Error(fiber.StatusBadRequest, i18n.ErrInvalidPassword, err))
}
