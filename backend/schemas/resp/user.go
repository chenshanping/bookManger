package resp

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username"` // 账号
	Sex      string `json:"sex"`
	Age      int    `json:"age"`
	Phone    string `json:"phone"`
}
