package req

import "gorm.io/gorm"

type BookCreate struct {
	gorm.Model
	Name    string `json:"name" binding:"required"`
	System  string `json:"system"  binding:"required"`
	Number  int32  `json:"number" binding:"required"`
	Author  string `json:"author" binding:"required"`
	Publish string `json:"publish" binding:"required"`
	Edition string `json:"edition" binding:"required"`
}
