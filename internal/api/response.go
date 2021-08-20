package api

import (
	"strings"

	"github.com/gofiber/fiber/v2"

	"go.vixal.xyz/esp/internal/i18n"
)

type Response struct {
	Code    int         `json:"code"`
	Err     string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Details string      `json:"details,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func (r Response) String() string {
	if r.Err != "" {
		return r.Err
	} else {
		return r.Message
	}
}

func (r Response) LowerString() string {
	return strings.ToLower(r.String())
}

func (r Response) Error() string {
	return r.Err
}

func (r Response) Success() bool {
	return r.Err == "" && r.Code < 400
}

func NewResponse(code int, id i18n.Message, params ...interface{}) Response {
	if code < fiber.StatusBadRequest {
		return Response{Code: code, Message: i18n.Msg(id, params...)}
	} else {
		return Response{Code: code, Err: i18n.Msg(id, params...)}
	}
}

func NewResponseWithData(code int, data interface{}, id i18n.Message, params ...interface{}) Response {
	if code < fiber.StatusBadRequest {
		return Response{Code: code, Message: i18n.Msg(id, params...), Data: data}
	} else {
		return Response{Code: code, Err: i18n.Msg(id, params...), Data: data}
	}
}

func NewResponseWithDataOrError(code int, data interface{}, err error, id i18n.Message, params ...interface{}) Response {
	var resp Response
	if code < fiber.StatusBadRequest {
		resp = Response{Code: code, Message: i18n.Msg(id, params...), Data: data}
	} else {
		resp = Response{Code: code, Err: i18n.Msg(id, params...), Data: data}
	}
	if err != nil {
		resp.Details = err.Error()
	}
	return resp
}
