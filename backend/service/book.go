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

func (bs bookservice) GetBookSeries(name string) (s model.Sort) {
	global.DB.Where("name=?", name).Find(&s)
	return
}

func (bs bookservice) CreateBook() (s model.Book) {
	global.DB.Create(&s)
	return
}
