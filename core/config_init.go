package core

import (
	"bookManage/global"
	"fmt"
	"github.com/spf13/viper"
)

func ConfigInit() {
	viper.SetConfigName("config") //找寻文件的名字
	viper.SetConfigType("yaml")   // 找寻文件的类型
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() //读取配置文件
	if err != nil {
		if v, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println(v)
		} else {
			panic(fmt.Errorf("read config err=%s", err))
		}
	}
	LoadMysqlConfig()
	LoadMioinConfig()
	LoadZapConfig()
	fmt.Println("配置文件初始化成功")
}

func LoadMysqlConfig() {
	global.Config.Mysql.Host = viper.GetString("mysql.host")
	global.Config.Mysql.User = viper.GetString("mysql.user")
	global.Config.Mysql.Password = viper.GetString("mysql.password")
	global.Config.Mysql.Port = viper.GetInt("mysql.port")
	global.Config.Mysql.Dbname = viper.GetString("mysql.dbname")

}

func LoadMioinConfig() {
	global.Config.Minio.Host = viper.GetString("minio.host")
	global.Config.Minio.User = viper.GetString("minio.user")
	global.Config.Minio.Password = viper.GetString("minio.password")
	global.Config.Minio.Port = viper.GetInt("minio.port")

}

func LoadZapConfig() {
	global.Config.Zap.Level = viper.GetString("zap.Level")
	global.Config.Zap.Format = viper.GetString("zap.Format")
	global.Config.Zap.Prefix = viper.GetString("zap.Prefix")
	global.Config.Zap.Director = viper.GetString("zap.Director")
	global.Config.Zap.LinkName = viper.GetString("zap.LinkName")
	global.Config.Zap.ShowLine = viper.GetBool("zap.ShowLine")
	global.Config.Zap.EncodeLevel = viper.GetString("zap.EncodeLevel")
	global.Config.Zap.StacktraceKey = viper.GetString("zap.StacktraceKey")
	global.Config.Zap.LogInConsole = viper.GetBool("zap.LogInConsole")

}
