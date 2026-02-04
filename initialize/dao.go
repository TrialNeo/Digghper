package initialize

import (
	"Diggpher/global"
	"Diggpher/internal/dao"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
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
		panic(err)
	}
	global.DataBase = DataBase
	dao.BindDao()
}
