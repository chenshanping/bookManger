package core

import (
	"bookManage/global"
	"bookManage/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func LogFormat(param gin.LogFormatterParams) string {
	// your custom format
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC3339Nano),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}
func GormInit() {
	host := global.Config.Mysql.Host
	user := global.Config.Mysql.User
	password := global.Config.Mysql.Password
	port := global.Config.Mysql.Port
	dbname := global.Config.Mysql.Dbname
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,  // 慢 SQL 阈值
			LogLevel:      logger.Error, // Log level
			Colorful:      true,         // 彩色打印
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("gorm初始化成功")
	}
	global.DB = db
	if err := global.DB.AutoMigrate(models.User{}, models.Book{}); err != nil {
		fmt.Printf("数据库创建失败:", err)
	}

}
