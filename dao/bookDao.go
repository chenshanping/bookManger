package dao

import (
	"bookManage/global"
	"bookManage/models"
)

func GetBookData(name string) models.Book {
	var book models.Book
	global.DB.Where("name", name).Find(&book)
	return book
}
func GetBookID(name string) int64 {
	var book models.Book
	global.DB.Where("name", name).Find(&book)
	return book.Id
}
func GetBookList() []models.Book {
	var books []models.Book
	global.DB.Find(&books)
	return books
}

func CreateBookDao(p *models.Book) {
	global.DB.Create(p)
}

func DeleteBookDao(bookname string) {
	var book models.Book
	id := GetBookID(bookname)
	global.DB.Where("id=?", id).Delete(&book)
}
func UpdateBook(p *models.Book) {
	global.DB.Model(p).Update("name", "")
	global.DB.Model(p).Update("dest", "")
}
