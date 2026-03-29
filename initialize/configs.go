package initialize

import (
	"Diggpher/global"
	"Diggpher/pkg/logger"
	"fmt"
	"github.com/spf13/viper"
)

func LoadConfigs() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./configs")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	err = viper.Unmarshal(global.CONFIG)
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
		return
	}

	// 初始化日志
	logConfig := &logger.Config{
		Level:   global.CONFIG.Logger.Level,
		Console: global.CONFIG.Logger.Console,
		Dir:     global.CONFIG.Logger.Dir,
	}
	logger.InitLogger(logConfig)
	global.Log = logger.Log
}
