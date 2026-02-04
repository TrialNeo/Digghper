package controller

import (
	"Diggpher/internal/service/errMsg"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Code    uint        `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data,omitempty"`
}

// respSuc 成功
func respSuc(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(&Response{
		Code:    errMsg.SUCCESS,
		Message: "respSuc",
		Data:    data,
	})
}

// respFail 业务失败（仍返回 HTTP 200）
func respFail(c *fiber.Ctx, code uint, message string) error {
	return c.Status(fiber.StatusOK).JSON(&Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}
