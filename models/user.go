package models

import "time"

type User struct {
	ID        uint   `gorm:"primarykey"`
	Username  string `json:"username" gorm:"not null" `
	Password  string `json:"password" gorm:"not null"`
	Token     string `json:"token"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName 表名默认会添加s,自定义表名
func (User) TableName() string {
	return "user"
}
