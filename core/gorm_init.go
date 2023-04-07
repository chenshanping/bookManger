package core

import (
	"bookManage/global"
	"bookManage/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func GormInit() {
	host := viper.GetString("mysql.host")
	user := viper.GetString("mysql.user")
	password := viper.GetString("mysql.password")
	port := viper.GetInt("mysql.port")
	dbname := viper.GetString("mysql.dbname")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)
	var mysqllogger logger.Interface
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqllogger,
	})
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("gorm初始化成功")
	}
	global.DB = db
	if err := global.DB.AutoMigrate(model.User{}, model.Book{}); err != nil {
		fmt.Printf("数据库创建失败:", err)
	}

}
