package router

import (
	"bookManage/controller"
	"bookManage/middleware"
	"github.com/gin-gonic/gin"
)

func LoadApiRouter(r *gin.Engine) {

	r.POST("/register", controller.RegisterHandler) //<<---这里将func指向controller下
	r.POST("/login", controller.LoginHandler)

	v1 := r.Group("/user")

	{
		v1.GET("/list", controller.GetUserList)
		v1.DELETE("/delete", controller.DeleteUser)

	}
	v2 := r.Group("/book")
	v2.Use(middleware.AuthMiddleware())
	{
		v2.GET("/list", controller.BookList)
		v2.POST("/create", controller.CreateBook)
		v2.DELETE("/delete", controller.DeleteBook)
		v2.POST("/updata", controller.UpdateBook)
		v2.GET("/bookdetail", controller.GetBookDetail)
	}

}
