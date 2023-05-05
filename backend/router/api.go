package router

import (
	"bookManage/controller"
	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	usergroup := r.Group("/user")
	//usergroup.Use(middleware.AuthMiddleware())
	{
		// 查看所有的user
		usergroup.GET("/list", controller.GetUserList)
		usergroup.DELETE("/delete/:id", controller.DeleteUser)
		usergroup.PUT("/update/:id", controller.UpdateUser)
		usergroup.GET("/list/:name", controller.GetUserDatail)
		usergroup.POST("/", controller.UserCreate)
	}
}
