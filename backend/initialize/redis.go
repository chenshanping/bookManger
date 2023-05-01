package initialize

import (
	"bookManage/global"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
)

func InitRedis() {
	redisconfig := global.Config.Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisconfig.Host, redisconfig.Port),
		Password: redisconfig.Password,
		DB:       redisconfig.DB,
	})
	timeoutCtx, cancelFunc := context.WithTimeout(context.Background(), redisconfig.DialTimeout)
	defer cancelFunc()
	_, err := redisClient.Ping(timeoutCtx).Result()
	if err != nil {
		panic(err)
	}
	global.Redis = redisClient
	log.Println("redis初始化成功")
}
