package route

import "Diggpher/global"

// BindRoute 对外暴露，绑定路由，以后换框架改这个位置
func BindRoute() {
	api := global.WebApp.Group("/api")
	bindAdminRoute(api)
}
