package main

import "Diggpher/initialize"

func main() {
	initialize.LoadConfigs()
	initialize.ConnRedis()
	initialize.ConnectDB()
	initialize.RunWebService()
}
