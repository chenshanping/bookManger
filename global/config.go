package global

import (
	"bookManage/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var DB *gorm.DB
var ZapLog *zap.Logger

var Config config.ConfigEnv
