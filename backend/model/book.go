package model

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Name    string `gorm:"column:name;default:NULL"`
	Sid     int32  `gorm:"column:sid;default:NULL"`
	Number  int32  `gorm:"column:number;default:NULL"`
	Author  string `gorm:"column:author;default:NULL"`
	Publish string `gorm:"column:publish;default:NULL"`
	Edition string `gorm:"column:edition;default:NULL"`
}

func (b *Book) TableName() string {
	return "book"
}
