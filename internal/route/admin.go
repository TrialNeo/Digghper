package route

import (
	"Diggpher/internal/controller"
	"Diggpher/internal/service"
	"Diggpher/pkg/middleware/auth"
	"github.com/gofiber/fiber/v2"
)

// bindAdminRoute 绑定admin的路由组
func bindAdminRoute(admin fiber.Router) {
	var (
		adminCtrl = &controller.AdminController{
			Service: new(service.AdminService),
		}
	)
	admin.Post("/admin/login", adminCtrl.Login)
	admin.Use(auth.MiddlewareAuth()).Route("/admin", func(router fiber.Router) {
		// 这是授权之后的使用hh
		router.Get("/hello", func(c *fiber.Ctx) error {
			return c.JSON(fiber.Map{
				"hello": "world",
			})
		})
	})
}
