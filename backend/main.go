package main

import (
	"bookManage/initialize"
)

func main() {
	initialize.InitConfig()
	initialize.InitDB()

	initialize.InitLogger()
	initialize.InitRedis()

	initialize.InitGin()
	//var userList []resp.User
	//var total int64
	//global.DB.Model(userList).Count(&total).Limit(1).Offset(1).Find(&userList)
	//fmt.Printf("%#v", userList)

}
