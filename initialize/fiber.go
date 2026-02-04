package initialize

import (
	"Diggpher/global"
	"Diggpher/internal/route"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func RunWebService() {
	global.WebApp = fiber.New(global.FbConfig)
	route.BindRoute()
	err := global.WebApp.Listen(fmt.Sprintf(":%d", global.CONFIG.Web.Port))
	if err != nil {
		panic(err.Error())
	}
}
