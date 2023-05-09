package service

import (
	"bookManage/global"
	"bookManage/model"
)

var BookService = bookservice{}

type bookservice struct {
}

func (bs bookservice) GetBookList(pageSize int, offSetval int) (bookList []model.Book, total int64) {
	global.DB.Model(&bookList).Count(&total).Limit(pageSize).Offset(offSetval).Find(&bookList)
	return
}
func (bs bookservice) GetBookID(name string) uint {
	var book model.Book
	global.DB.Where("name=?", name).Find(&book)
	return book.ID
}

func (bs bookservice) CreateBook(s model.Book) (err error) {
	if err := global.DB.Create(&s).Error; err != nil {
		return err
	}
	return
}
