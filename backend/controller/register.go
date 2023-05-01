package controller

import (
	"bookManage/model"
	"bookManage/schemas/resp"
	"bookManage/service"
	"bookManage/util"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var user model.User
	if err := c.BindJSON(&user); err != nil {
		resp.FailWithMsg(c, resp.ParamsValidError, err.Error())
		return
	}
	if service.UserService.GetUserId(user.Username) != 0 {
		resp.Fail(c, resp.RegisterError)
		return
	}
	user.Password = util.ToolsUtil.MakeMd5(user.Password)
	if err := service.UserService.CreateUser(&user); err != nil {
		resp.FailWithMsg(c, resp.Failed, err.Error())
	} else {
		resp.OkWithMsg(c, "注册成功")
	}

}
