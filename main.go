package main

import "bookManage/core"

func main() {

	core.ConfigInit()

	//gorm初始化
	core.GormInit()

	// swagger
	core.Swagger()
	//路由初始化
	r := core.RouterInit()
	err := r.Run(":8000")
	if err != nil {
		return
	}

}
