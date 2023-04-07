package router

import (
	"bookManage/controller"
	"github.com/gin-gonic/gin"
)

func LoadApiRouter(r *gin.Engine) {
	r.POST("/register", controller.RegisterHandler) //<<---这里将func指向controller下
	r.GET("/login", controller.LoginHandler)

	v1 := r.Group("/api/v1")
	v1.GET("/user", controller.GetUserList)
}
