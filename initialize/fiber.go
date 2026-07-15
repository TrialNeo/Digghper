package initialize

import (
	"Diggpher/global"
	"Diggpher/internal/route"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func RunWebService() {
	global.WebApp = fiber.New(global.FbConfig)
	global.WebApp.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3007,http://10.150.83.58:3007",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Content-Type,Authorization,X-Requested-With",
	}))
	route.BindRoute()
	err := global.WebApp.Listen(fmt.Sprintf(":%d", global.CONFIG.Web.Port))
	if err != nil {
		panic(err.Error())
	}
}
