package dao

import (
	"bookManage/global"
	"bookManage/models"
	"fmt"
	"log"
)

func MysqlAuto() {
	if err := global.DB.AutoMigrate(models.User{}, models.Book{}, models.BookUser{}); err != nil {
		log.Fatalln(fmt.Sprintf("数据库创建失败:%s", err))
	}
}
