package controller

import (
	"Diggpher/internal/service/errMsg"

	"github.com/gofiber/fiber/v2"
)

type Response[T any] struct {
	Code    uint        `json:"code"`
	Message string      `json:"msg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type RespondIMP struct {
	c    *fiber.Ctx
	code uint
}

func newRespondIMP(c *fiber.Ctx) *RespondIMP {
	return &RespondIMP{c: c}
}

func (r *RespondIMP) withCode(code uint) *RespondIMP {
	r.code = code
	return r
}

func (r *RespondIMP) Respond(data any) error {
	if r.code != errMsg.SUCCESS {
		return r.c.Status(fiber.StatusOK).JSON(&Response[any]{
			Code:    r.code,
			Message: errMsg.GetErrMsg(r.code),
		})
	}
	return r.c.Status(fiber.StatusOK).JSON(&Response[any]{
		Code:    r.code,
		Message: errMsg.GetErrMsg(r.code),
		Data:    data,
	})
}

// 旧版兼容
func Respond(c *fiber.Ctx, code uint, data interface{}) error {
	r := newRespondIMP(c)
	return r.withCode(code).Respond(data)
}
