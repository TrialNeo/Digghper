package initialize

import (
	"Diggpher/global"
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
}
