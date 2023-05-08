package initialize

import (
	"bookManage/global"
	"bookManage/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// 设置新的Logger
func setNewLogger(gConfig *gorm.Config) {
	m := global.Config.Mysql
	logPath := global.Config.Log.Path
	tStr := time.Now().Format("20060102")
	sqlfile := fmt.Sprintf("/sql-%s.log", tStr)
	file, _ := os.OpenFile(logPath+sqlfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	// 日志级别映射 error、info、warn

	logLevelMap := map[string]logger.LogLevel{

		"error": logger.Error,
		"info":  logger.Info,
		"warn":  logger.Warn,
	}
	var logLevel logger.LogLevel
	var ok bool
	if logLevel, ok = logLevelMap[m.LogLevel]; !ok {
		logLevel = logger.Error
	}
	newLogger := logger.New(log.New(file, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             m.SlowSql,                   //慢SQL时间
		LogLevel:                  logLevel,                    // 记录日志级别
		IgnoreRecordNotFoundError: m.IgnoreRecordNotFoundError, // 是否忽略ErrRecordNotFound(未查到记录错误)
		Colorful:                  false,                       // 开关颜色
	})
	gConfig.Logger = newLogger
}
func test() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 彩色打印
		},
	)
	return newLogger

}
func InitDB() {
	m := global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%v&loc=%s",
		m.User,
		m.Password,
		m.Host,
		m.Port,
		m.Dbname,
		m.CharSet,
		m.ParseTime,
		m.TimeZone,
	)
	gormConfig := &gorm.Config{
		Logger: test(),
		//是否跳过默认事务
		SkipDefaultTransaction: m.Gorm.SkipDefaultTx,
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   m.Gorm.TablePrefix,
			SingularTable: m.Gorm.SingularTable,
		},
		// 执行任何SQL时都会创建一个prepared statement并将其缓存，以提高后续的效率
		PrepareStmt: m.Gorm.PrepareStmt,
		//在AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束，若要禁用该特性，可将其设置为 true
		DisableForeignKeyConstraintWhenMigrating: m.Gorm.CloseForeignKey,
	}
	// 是否覆盖默认sql配置
	system := global.Config.System
	if system.Env != "debug" && m.Gorm.CoverLogger {
		setNewLogger(gormConfig)
	}
	mysqlConfig := mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         255,
		SkipInitializeWithVersion: false,
	})
	db, err := gorm.Open(mysqlConfig, gormConfig)
	if err != nil {
		panic("failed to connect database")
	}
	global.DB = db
	MysqlAuto()
	log.Println("初始化数据库成功")

}
func MysqlAuto() {
	if err := global.DB.AutoMigrate(model.User{}, model.Book{}); err != nil {
		log.Fatalln(fmt.Sprintf("数据库创建失败:%s", err))
	}
}
