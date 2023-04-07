package model

type User struct {
	Id       int64  `json:"id gorm:primaryKey"`
	Username string `json:"username" gorm:"not null" binding:"required"`
	Password string `json:"password" gorm:"not null" binding:"required"`
	Token    string `json:"token"`
}

// TableName 表名默认会添加s,自定义表名
func (User) TableName() string {
	return "user"
}
