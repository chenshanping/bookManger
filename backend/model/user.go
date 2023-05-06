package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required,min=2,max=20"` // 账号
	// gorm:"default 123456"
	Password string `json:"password" binding:"required,min=6,max=32" ` // 密码
	Sex      string `json:"sex" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Phone    string `json:"phone" binding:"required,min=9,max=11"`
}

func (User) TableName() string {
	return "users"
}
