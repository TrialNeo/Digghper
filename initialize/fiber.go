package initialize

import (
	"Diggpher/global"
	"Diggpher/internal/route"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func RunWebService() {
	global.Log.Info("Starting web service")
	global.WebApp = fiber.New(global.FbConfig)
	route.BindRoute()
	global.Log.Info("Routes bound successfully")
	err := global.WebApp.Listen(fmt.Sprintf(":%d", global.CONFIG.Web.Port))
	if err != nil {
		global.Log.Fatal("Failed to start web service", zap.Error(err))
	}
	global.Log.Info("Web service started successfully")
}
