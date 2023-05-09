package router

import (
	"bookManage/controller"
	"github.com/gin-gonic/gin"
)

func SetRouter(r *gin.Engine) {
	r.POST("/register", controller.Register)
	r.POST("/login", controller.Login)
	api := r.Group("/api")
	usergroup := api.Group("/user")
	//usergroup.Use(middleware.AuthMiddleware())
	{
		// 查看所有的user
		usergroup.GET("/list", controller.GetUserList)
		usergroup.DELETE("/delete/:id", controller.DeleteUser)
		usergroup.PUT("/update/:id", controller.UpdateUser)
		usergroup.GET("/list/:name", controller.GetUserDatail)
		usergroup.POST("/", controller.UserCreate)
	}
	bookgroup := api.Group("/book")
	{
		bookgroup.GET("/list", controller.BookApi.GetBookList)
		bookgroup.POST("/add", controller.BookApi.CreateBook)

	}
	sortgroup := api.Group("/sort")
	{
		sortgroup.GET("/list", controller.SortApi.GetSortList)
		sortgroup.GET(":id", controller.SortApi.GetSortName)
	}
}
