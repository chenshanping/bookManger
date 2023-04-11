package response

import (
	"github.com/gin-gonic/gin"
)

//type ResponseData struct {
//	Code ResCode     `json:"code"`
//	Data interface{} `json:"data"`
//	Msg  interface{} `json:"msg"`
//}

var (
	Success      = 200
	InvalidParam = 400
	AuthParm     = 401
)

func ResponWrite(c *gin.Context, msg string) {
	c.JSON(Success, gin.H{
		"code":   1,
		"status": Success,
		"msg":    msg,
	})
}
func ResponSuccess(c *gin.Context, data interface{}) {
	c.JSON(Success, gin.H{
		"code": 1,
		"msg":  "请求成功",
		"data": data,
	})
}

func ResponParam(c *gin.Context) {
	c.JSON(InvalidParam, gin.H{
		"code":   0,
		"status": InvalidParam,
		"msg":    "请求参数无效",
	})
}
func ResponAuth(c *gin.Context) {
	c.JSON(AuthParm, gin.H{
		"code":   0,
		"status": AuthParm,
		"msg":    "账户或者密码错误",
	})

}
