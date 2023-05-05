package controller

import (
	"bookManage/model"
	"bookManage/schemas/req"
	"bookManage/schemas/resp"
	"bookManage/service"
	"bookManage/util"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUserList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	offsetval := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offsetval = -1
	}
	list, total := service.UserService.GetAllUser(pageSize, offsetval)
	if len(list) == 0 {
		resp.FailWithMsg(c, resp.Failed, "没有查询到数据")
	} else {
		h := gin.H{
			"list":     list,
			"total":    total,
			"pageNum":  pageNum,
			"pageSize": pageSize,
		}
		resp.OkWithMsgData(c, "查询成功", h)
	}

}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	exist := service.UserService.IDExist(id)
	if exist == false {
		resp.FailWithMsg(c, resp.Failed, "ID不存在,删除失败")
	} else {
		service.UserService.DeleteUser(id)
		resp.OkWithMsg(c, "删除成功")
	}

}

func UpdateUser(c *gin.Context) {
	var user model.User
	id := c.Param("id")
	if err := c.BindJSON(&user); err != nil {
		resp.FailWithMsg(c, resp.ParamsValidError, err.Error())
		return
	}
	exist := service.UserService.IDExist(id)
	if exist == false {
		resp.FailWithMsg(c, resp.Failed, "ID不存在,无法修改")
	} else {
		user.Password = util.ToolsUtil.MakeMd5(user.Password)
		service.UserService.UpdateUser(id, user)
		resp.OkWithMsg(c, "修改成功")
	}

}

func GetUserDatail(c *gin.Context) {
	username := c.Param("name")
	detail := service.UserService.UserDetail(username)
	if len(detail) == 0 {
		resp.FailWithMsg(c, resp.Failed, "没有查询到数据")
	} else {
		resp.OkWithMsgData(c, "查询成功", detail)
	}

}

func UserCreate(c *gin.Context) {
	var user req.UserCreate
	if err := c.BindJSON(&user); err != nil {
		resp.FailWithMsg(c, resp.ParamsValidError, err.Error())
		return
	}
	if service.UserService.GetUserId(user.Username) != 0 {
		resp.Fail(c, resp.RegisterError)
		return
	}
	if err := service.UserService.CreateUser(&user); err != nil {
		resp.FailWithMsg(c, resp.Failed, err.Error())
	} else {
		resp.OkWithMsg(c, "用户创建成功")
	}
}
