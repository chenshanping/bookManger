package controller

import (
	"bookManage/global"
	"bookManage/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

//func ifUsername(user string) {
//	m := model.User{}
//
//}

// RegisterHandler 注册功能
func RegisterHandler(c *gin.Context) {
	p := new(model.User)
	// 参数校验和绑定
	if err := c.BindJSON(p); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	// 入库
	u := model.User{Username: p.Username}
	tx := global.DB.Where("username", p.Username).First(&u)
	c.JSON(200, tx)
	//global.DB.Create(p)
	//c.JSON(200, gin.H{"msg": p})
}

func LoginHandler(c *gin.Context) {
	p := new(model.User)
	// 参数校验和绑定
	if err := c.BindJSON(p); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	u := model.User{Username: p.Username, Password: p.Password}
	// 判断rows是否有数据
	if rows := global.DB.Where(&u).First(&u).Row(); rows == nil {
		c.JSON(403, gin.H{"msg": "用户名密码错误"})
		return
	}
	//随机生成一个字符串作为Token
	token := uuid.New().String()
	global.DB.Model(&u).Update("token", token)
	c.JSON(200, gin.H{"msg": "success"})
}

func GetUserList(c *gin.Context) {
	var users []model.User

	global.DB.Find(&users)
	c.JSON(200, gin.H{"user": users})
}
