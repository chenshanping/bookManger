package req

import "gorm.io/gorm"

type UserLoginReq struct {
	Username string `json:"username" binding:"required,min=2,max=20"` // 账号
	Password string `json:"password" binding:"required,min=6,max=32"` // 密码
}

func (UserLoginReq) TableName() string {
	return "book_user"
}

type UserRegister struct {
	gorm.Model
	Username string `json:"username" binding:"required,min=2,max=20"` // 账号
	Password string `json:"password" binding:"required,min=6,max=32"` // 密码
	Sex      string `json:"sex" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Phone    string `json:"phone" binding:"required,min=9,max=11"`
}

func (UserRegister) TableName() string {
	return "book_user"
}

type UserCreate struct {
	Username string `json:"username" binding:"required,min=2,max=20"` // 账号
	Sex      string `json:"sex" binding:"required"`
	Age      int    `json:"age" binding:"required"`
	Phone    string `json:"phone" binding:"required,min=9,max=11"`
}

func (UserCreate) TableName() string {
	return "book_user"
}
