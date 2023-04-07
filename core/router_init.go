package core

import (
	"bookManage/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func RouterInit() *gin.Engine {
	// 1. 初始化
	//gin.SetMode("release")
	r := gin.Default()

	// 2. 定义路由
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Success")
	})
	router.TestApi(r)
	router.LoadApiRouter(r)
	fmt.Println("路由初始化成功")
	return r
}
