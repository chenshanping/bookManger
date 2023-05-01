package util

import (
	"bookManage/global"
	"golang.org/x/net/context"
	"time"
)

var RedisUtil = redisUtil{}

type redisUtil struct{}

func (ru redisUtil) DBSize() int64 {
	size, err := global.Redis.DBSize(context.Background()).Result()
	if err != nil {
		global.Logger.Errorf("redisUtil.DBSize err: err=[%+v]", err)
		return 0
	}
	return size
}

func (ru redisUtil) Get(key string) string {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	res, err := global.Redis.Get(ctx, key).Result()
	if err != nil {
		global.Logger.Errorf("redisUtil.Get err: err=[%+v]", err)
		return ""
	}
	return res
}
func (ru redisUtil) Set(key string, value interface{}, timeSec int) bool {
	err := global.Redis.Set(context.Background(),
		global.Config.Redis.Prefix+key, value, time.Duration(timeSec)*time.Second).Err()
	if err != nil {
		global.Logger.Errorf("redisUtil.Set err: err=[%+v]", err)
		return false
	}
	return true
}
