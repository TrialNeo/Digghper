package route

import (
	"Diggpher/internal/controller"
	"Diggpher/internal/service"
	"Diggpher/pkg/middleware/auth"

	"github.com/gofiber/fiber/v2"
)

type adminRouters struct {
	adminCtrl *controller.AdminController
}

// WithAdminRoute 绑定admin的路由组
func WithAdminRoute(admin fiber.Router) {
	a := &adminRouters{
		adminCtrl: &controller.AdminController{
			Service: service.NewAdminService(),
		},
	}
	admin.Post("/admin/login", a.adminCtrl.Login)
	admin.Use(auth.MiddlewareAuth()).Route("/admin", func(router fiber.Router) {
		// 授权之后的路由
		_ = router
	})
}
