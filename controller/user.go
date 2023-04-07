package controller

import (
	"bookManage/global"
	"bookManage/models"
	"bookManage/models/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"math/rand"
)

//func ifUsername(user string) {
//	m := models.User{}
//
//}

// RegisterHandler 注册功能
func RegisterHandler(c *gin.Context) {
	var user models.User

	// 参数校验和绑定
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"msg": err.Error()})
		return
	}
	// 入库
	fmt.Printf("%s %s", user.Username, user.Password)
	global.DB.Create(&user)
	//tx := global.DB.Create(user)
	//if tx.Error != nil {
	//	fmt.Printf("%s", tx.Error)
	//	return
	//}
	c.JSON(200, gin.H{"msg": user})
	//global.DB.Where("username=?", user.Username).First(&user)
	//if user.Password == "" {
	//
	//
	//} else {
	//
	//}
	//global.DB.Create(p)
	//c.JSON(200, gin.H{"msg": p})
}

func LoginHandler(c *gin.Context) {
	p := new(models.User)

	// 参数校验和绑定
	if err := c.ShouldBind(p); err != nil {
		response.ResponseError(c, response.CodeNoParam)
		return
	}
	if p.Username == "" || p.Password == "" {
		response.ResponseError(c, response.CodeInvalidParam)
		return
	}
	// 判断用户名密码是否正确
	u := models.User{Username: p.Username, Password: p.Password}
	// 判断rows是否有数据
	if rows := global.DB.Where(&u).First(&u).Row(); rows == nil {
		fmt.Printf("%v", rows)
		response.ResponseError(c, response.CodeInvalidPassword)
		//c.JSON(403, gin.H{"msg": "用户名密码错误"})
		return
	}
	//随机生成一个字符串作为Token
	token := uuid.New().String()
	global.DB.Model(&u).Update("token", token)

	response.ResponseSuccess(c, nil)
	//c.JSON(200, gin.H{"msg": "success"})
}

func GetUserList(c *gin.Context) {
	//c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	var users []models.User
	//var list1 [2][2]string
	err := global.DB.Find(&users).Error
	if err != nil {
		return
	}

	return

}
func DeleteUser(c *gin.Context) {
	value := c.Query("username")

	user := models.User{}
	global.DB.Where("username=?", value).Delete(&user)
	if user.Password != "" {
		c.JSON(200, gin.H{
			"msg": fmt.Sprintf("用户%s删除成功", value),
		})
	} else {
		c.JSON(200, gin.H{
			"msg": fmt.Sprintf("用户%s已经删除", user.Username),
		})
	}

}

func QQ(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	legendData := []string{"周一", "周二", "周三", "周四", "周五", "周六", "周日"}
	xAxisData := []int{120, 240, rand.Intn(500), rand.Intn(500), 150, 230, 180}
	ginSwagger.DisablingWrapHandler(swaggerfiles.Handler, "SWAGGER")
	c.JSON(200, gin.H{
		"legend_data": legendData,
		"xAxis_data":  xAxisData,
	})

}
