package router

import (
	"bookManage/controller"
	"github.com/gin-gonic/gin"
)

func LoadApiRouter(r *gin.Engine) {

	r.POST("/register", controller.RegisterHandler) //<<---这里将func指向controller下
	r.POST("/login", controller.LoginHandler)

	v1 := r.Group("/user")
	{
		v1.GET("/list", controller.GetUserList)
		v1.GET("/delete", controller.DeleteUser)

	}
	group := r.Group("/v1")
	{
		group.GET("/line", controller.QQ)
	}

}
