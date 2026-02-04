package controller

import (
	"Diggpher/internal/service"
	"Diggpher/internal/service/errMsg"
	"github.com/gofiber/fiber/v2"
)

type AdminController struct {
	Service *service.AdminService
}

// Login 登录
func (a *AdminController) Login(c *fiber.Ctx) error {
	login := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	if err := c.BodyParser(&login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": err.Error()})
	}
	resp := a.Service.Login(login.Username, login.Password, c.IP())
	if resp.Code == errMsg.SUCCESS {
		return respSuc(c, fiber.Map{
			"token": resp.Token,
		})
	} else {
		return respFail(c, resp.Code, resp.ErrMsg)
	}
}
