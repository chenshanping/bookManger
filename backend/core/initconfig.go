package core

import (
	"bookManage/global"
	"fmt"
	"github.com/spf13/viper"
	"log"
)

func InitConfig() {
	viper.SetConfigName("config") //找寻文件的名字
	viper.SetConfigType("yaml")   // 找寻文件的类型
	viper.AddConfigPath("backend")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Printf("配置文件未找到！%v\n", err)
			return
		} else {
			panic(fmt.Sprintf("找到配置文件,但是解析错误,%v\n", err))
		}
	}
	ConfigStruct()
	log.Print("配置文件初始化成功")
}

func ConfigStruct() {
	// 映射到结构体
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("配置映射错误,%v\n", err)
	}

}
