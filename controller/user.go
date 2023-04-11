package controller

import (
	"bookManage/dao"
	"bookManage/models"
	"bookManage/models/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

// RegisterHandler 注册功能
func RegisterHandler(c *gin.Context) {
	var user models.User
	// 参数校验和绑定
	if err := c.BindJSON(&user); err != nil {
		response.ResponParam(c)
		return
	}
	// 检验参数是否写错了
	if user.Username == "" || user.Password == "" {
		response.ResponParam(c)
		return
	}
	if dao.GetUserid(user.Username) != 0 {
		response.ResponWrite(c, fmt.Sprintf("%s用名名已经注册了", user.Username))
		return
	}
	//dao.UserDao().PassWordCrypto()
	user.Password = dao.PassWordCrypto(user.Password)
	dao.CreateUser(&user)
	response.ResponWrite(c, "创建成功")

}

// LoginHandler 登录功能
func LoginHandler(c *gin.Context) {
	/*
		提交数据 转为md5 与数据库对比

	*/
	var user models.User

	// 参数校验和绑定
	if err := c.ShouldBind(&user); err != nil {
		response.ResponParam(c)
		return
	}
	if user.Username == "" || user.Password == "" {
		response.ResponParam(c)
		return
	}
	// 判断用户名密码是否正确
	// 判断rows是否有数据
	requestPwd := dao.PassWordCrypto(user.Password)
	mysqlPwd := dao.GetUserData(user.Username).Password
	if requestPwd != mysqlPwd {
		response.ResponAuth(c)
		return
	}
	//if rows := global.DB.Where(&user).First(&user).Row(); rows == nil {
	//	response.ResponAuth(c)
	//	return
	//}
	//随机生成一个字符串作为Token
	dao.SetUserToken(&user)
	response.ResponWrite(c, "登录成功")
}

func GetUserList(c *gin.Context) {
	userlist := dao.GetUserList()
	response.ResponSuccess(c, userlist)

}
func DeleteUser(c *gin.Context) {
	value := c.Query("username")
	dao.DeleteUser(value)
	response.ResponWrite(c, "删除成功")

}
