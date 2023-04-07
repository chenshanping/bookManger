package core

import (
	"fmt"
	"github.com/spf13/viper"
)

func ConfigInit() {
	viper.SetConfigName("config") //找寻文件的名字
	viper.SetConfigType("yaml")   // 找寻文件的类型
	viper.AddConfigPath("conf")   //从当前目录下的哪个文件夹找寻，
	viper.AddConfigPath(".")      //.代表当前文件夹找寻，可以多个目录找寻，生成数组
	err := viper.ReadInConfig()   //读取配置文件
	if err != nil {
		if v, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println(v)
		} else {
			panic(fmt.Errorf("read config err=%s", err))
		}
	}
	fmt.Println("配置文件初始化成功")
}
