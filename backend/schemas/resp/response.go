package resp

import (
	"bookManage/global"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type RespType struct {
	code int
	msg  string
	data interface{}
}

// Response 响应格式结构
type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

var (
	Success       = RespType{code: 200, msg: "成功"}
	RegisterError = RespType{code: 200, msg: "账号已被注册"}
	Failed        = RespType{code: 200, msg: "失败"}

	ParamsValidError  = RespType{code: 400, msg: "参数校验错误"}
	LoginAccountError = RespType{code: 401, msg: "账号或密码错误"}

	TokenEmpty   = RespType{code: 401, msg: "token参数为空"}
	TokenInvalid = RespType{code: 401, msg: "token参数无效"}
	NoPermission = RespType{code: 403, msg: "无相关权限"}

	SystemError = RespType{code: 500, msg: "系统错误"}
)

func Result(c *gin.Context, resp RespType, data interface{}) {
	if data == nil {
		data = resp.data
	}
	c.JSON(http.StatusOK, Response{
		Code: resp.code,
		Msg:  resp.msg,
		Data: data,
	})
}

// Ok 正常响应
func Ok(c *gin.Context) {
	Result(c, Success, []string{})
}

func OkWithData(c *gin.Context, data interface{}) {
	Result(c, Success, data)
}

func OkWithMsg(c *gin.Context, msg string) {
	resp := Success
	resp.msg = msg
	Result(c, resp, []string{})
}
func OkWithMsgData(c *gin.Context, msg string, data interface{}) {
	resp := Success
	resp.msg = msg
	Result(c, resp, data)
}
func respLogger(resp RespType, template string, args ...interface{}) {
	loggerFunc := global.Logger.WithOptions(zap.AddCallerSkip(2)).Warnf
	if resp.code >= 500 {
		loggerFunc = global.Logger.WithOptions(zap.AddCallerSkip(1)).Errorf
	}
	loggerFunc(template, args...)
}
func Fail(c *gin.Context, resp RespType) {
	respLogger(resp, "Request Fail: url=[%s], resp=[%+v]", c.Request.URL.Path, resp)
	Result(c, resp, []string{})
}

func FailWithMsg(c *gin.Context, resp RespType, msg string) {
	resp.msg = msg
	respLogger(resp, "Request FailWithMsg: url=[%s], resp=[%+v]", c.Request.URL.Path, resp)
	Result(c, resp, []string{})
}
