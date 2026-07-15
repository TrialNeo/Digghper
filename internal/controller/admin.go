package controller

import (
	"Diggpher/internal/service"
	"Diggpher/internal/service/errMsg"

	"github.com/gofiber/fiber/v2"
)

type AdminController struct {
	Service *service.AdminService
}

func (a *AdminController) Login(c *fiber.Ctx) error {
	re := newRespondIMP(c)
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&req); err != nil {
		return re.withCode(errMsg.ErrorInvalidParams).Respond(nil)
	}
	resp := a.Service.Login(req.Username, req.Password, c.IP())
	return re.withCode(resp.Code).Respond(fiber.Map{
		"token": resp.Token,
	})
}
