package initialize

import (
	"Diggpher/global"
	"Diggpher/internal/dao"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	global.Log.Info("Connecting to database")
	DataBase, err := gorm.Open(postgres.Open(
		fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
			global.CONFIG.Database.Host,
			global.CONFIG.Database.User,
			global.CONFIG.Database.Psw,
			global.CONFIG.Database.DataSourceName,
			global.CONFIG.Database.Port,
			global.CONFIG.Database.TimeZone,
		)),
		&gorm.Config{})
	if err != nil {
		global.Log.Fatal("Failed to connect to database", zap.Error(err))
	}
	global.DataBase = DataBase
	global.Log.Info("Database connected successfully")
	dao.BindDao()
	global.Log.Info("Database tables bound successfully")
}
