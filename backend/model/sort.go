package model

type Sort struct {
	Id   int32  `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	Name string `gorm:"column:name;default:NULL"`
}

func (s *Sort) TableName() string {
	return "sort"
}
