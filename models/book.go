package models

type Book struct {
	Id    int64  `json:"id" gorm:"primaryKey"`
	Name  string `json:"name" gorm:"not null" binding:"required"`
	Desc  string `json:"desc"`
	Users []User `gorm:"many2many:book_users"`
}

func (Book) TableName() string {
	return "book"
}