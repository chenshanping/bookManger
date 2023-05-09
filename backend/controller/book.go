package controller

import (
	"bookManage/model"
	"bookManage/schemas/resp"
	"bookManage/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

var BookApi = bookApi{}

type bookApi struct {
}

func (ba bookApi) GetBookList(c *gin.Context) {
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

func (ba bookApi) CreateBook(c *gin.Context) {
	var book model.Book
	if err := c.BindJSON(&book); err != nil {
		resp.FailWithMsg(c, resp.ParamsValidError, err.Error())
		return
	}

	if service.BookService.GetBookID(book.Name) != 0 {
		resp.OkWithMsg(c, "这本图书已经存在")
	} else {
		err := service.BookService.CreateBook(book)
		if err != nil {
			resp.FailWithMsg(c, resp.Failed, err.Error())
		} else {
			resp.OkWithMsg(c, "图书创建成功")
		}

	}
}
