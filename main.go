package main

import "bookManage/core"

func main() {
	//配置文件初始化
	core.ConfigInit()
	//gorm初始化
	core.GormInit()
	//路由初始化
	router := core.RouterInit()
	router.Run(":8000")

}
