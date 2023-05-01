package controller

import (
	"bookManage/schemas/req"
	"bookManage/schemas/resp"
	"bookManage/service"
	"bookManage/util"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var ulr req.UserLoginReq
	if err := c.BindJSON(&ulr); err != nil {
		resp.FailWithMsg(c, resp.ParamsValidError, err.Error())
		return
	}
	md5 := util.ToolsUtil.MakeMd5(ulr.Password)
	u1 := service.UserService.FindByUsername(ulr.Username)
	if u1.Password == md5 {
		key := fmt.Sprintf("user_%s", ulr.Username)
		set := util.RedisUtil.Set(util.ToolsUtil.MakeUuid(), key, 7200)
		if set {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}

		resp.OkWithMsg(c, "登录成功")
	} else {
		resp.Fail(c, resp.LoginAccountError)
	}

}
