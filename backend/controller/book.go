package controller

import (
	"bookManage/schemas/resp"
	"bookManage/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetSid(c *gin.Context) {
	name := c.Query("name")
	s := service.BookService.GetBookSeries(name)
	if s.Id == 0 {
		resp.OkWithMsg(c, "没有查询到数据")
		return
	}
	resp.OkWithMsgData(c, "查询成功", s)
}

func GetBookList(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))

	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	offSetval := (pageNum - 1) * pageSize
	if pageNum == -1 && pageSize == -1 {
		offSetval = -1
	}
	list, total := service.BookService.GetBookList(pageSize, offSetval)
	data := gin.H{
		"total":    total,
		"pageNum":  offSetval,
		"pageSize": pageSize,
		"list":     list,
	}
	if len(list) == 0 {
		resp.OkWithMsg(c, "没有查询到数据")
	} else {
		resp.OkWithMsgData(c, "查询成功", data)
	}

}
