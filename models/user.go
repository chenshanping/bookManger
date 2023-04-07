package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"not null" `
	Password string `json:"password" gorm:"not null"`
	Token    string `json:"token"`
}

// TableName 表名默认会添加s,自定义表名
func (User) TableName() string {
	return "user"
}
