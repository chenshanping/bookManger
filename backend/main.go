package main

import "bookManage/initialize"

func main() {
	initialize.InitConfig()
	initialize.InitDB()

	initialize.InitLogger()
	initialize.InitRedis()

	initialize.InitGin()

}
