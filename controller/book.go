package controller

import (
	"bookManage/dao"
	"bookManage/global"
	"bookManage/models"
	"bookManage/models/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func CreateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBind(&book); err != nil {
		response.ResponParam(c)
		return
	}
	if dao.GetBookID(book.Name) != 0 {
		response.ResponWrite(c, "这个书籍已经存在")
		return
	}
	dao.CreateBookDao(&book)
	response.ResponWrite(c, "创建书籍成功")

}

func BookList(c *gin.Context) {
	booklist := dao.GetBookList()
	response.ResponSuccess(c, booklist)
}
func DeleteBook(c *gin.Context) {
	value := c.Query("name")
	if value == "" {
		response.ResponParam(c)
		return
	}
	dao.DeleteBookDao(value)
	response.ResponWrite(c, "删除书籍成功")

}
func GetBookDetail(c *gin.Context) {
	value := c.Query("name")
	id := dao.GetBookID(value)
	if id == 0 {
		response.ResponWrite(c, "无此书籍")
		return
	}
	data := dao.GetBookData(value)
	response.ResponSuccess(c, data)
}
func UpdateBook(c *gin.Context) {
	var book models.Book
	if err := c.ShouldBind(&book); err != nil {
		response.ResponParam(c)
		return
	}
	id := dao.GetBookID(book.Name)
	if id == 0 {
		response.ResponWrite(c, "找不到这个本书籍无法更新")
		return
	}

	global.DB.Model(&book).Where("id=?", id).Update("name", book.Name)
	global.DB.Model(&book).Where("id=?", id).Update("desc", book.Desc)
	response.ResponWrite(c, fmt.Sprintf("%s书籍更新成功", book.Name))
}
