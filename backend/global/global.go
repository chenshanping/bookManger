package global

import (
	"bookManage/config"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var Config config.Config

var Logger *zap.SugaredLogger

var DB *gorm.DB
var Redis *redis.Client
