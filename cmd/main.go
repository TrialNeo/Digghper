package main

import (
	"Diggpher/global"
	"Diggpher/initialize"
	"Diggpher/pkg/logger"
)

func main() {
	// 先初始化默认日志，用于配置加载过程
	logger.InitLogger(logger.DefaultConfig())
	global.Log = logger.Log
	global.Log.Info("Starting application")
	initialize.LoadConfigs()
	initialize.ConnRedis()
	initialize.ConnectDB()
	initialize.RunWebService()
}
