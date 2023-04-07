package models

// BookUser 用户与书籍关联关系表
type BookUser struct {
	UserID int64 `gorm:"primaryKey"`
	BookID int64 `gorm:"primaryKey"`
}
