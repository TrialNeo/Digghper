package dao

import (
	"Diggpher/global"
)

// BindDao 对外接口，dao绑定
func BindDao() {
	//下面数据库初始化失败就不能继续用了。。
	err := global.DataBase.AutoMigrate(new(Admin))
	if err != nil {
		panic(err)
	}

}
