package initialize

import (
	"Diggpher/global"
	"Diggpher/internal/dao"
	"fmt"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB 连接数据库
func ConnectDB() {
	var (
		DataBase *gorm.DB
		err      error
	)
	global.Log.Info("Choose database", zap.String("name", global.CONFIG.Database.DriverName))
	switch global.CONFIG.Database.DriverName {
	case "mysql":
		DataBase, err = gorm.Open(mysql.Open(
			fmt.Sprintf(
				"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&sql_mode=NO_ENGINE_SUBSTITUTION",
				global.CONFIG.Database.User,
				global.CONFIG.Database.Psw,
				global.CONFIG.Database.Host,
				global.CONFIG.Database.Port,
				global.CONFIG.Database.DataSourceName,
			)),
			&gorm.Config{})
	case "postgres":
		DataBase, err = gorm.Open(postgres.Open(fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=%s",
			global.CONFIG.Database.Host,
			global.CONFIG.Database.User,
			global.CONFIG.Database.Psw,
			global.CONFIG.Database.DataSourceName,
			global.CONFIG.Database.Port,
			global.CONFIG.Database.TimeZone,
		)), &gorm.Config{})
	default:
		global.Log.Fatal("no such database type supported")
	}
	if err != nil {
		panic(err)
	}
	global.DataBase = DataBase
	dao.BindDao()
}
