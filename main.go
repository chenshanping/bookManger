package main

import (
	"bookManage/core"
	"bookManage/global"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func QQ(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start).Milliseconds()
		// 写入日志
		global.ZapLog.Info(fmt.Sprintf("[GIN] %s", path),
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("cost", fmt.Sprintf("%dms", cost)),
		)
	}
}
func main() {

	core.ConfigInit()

	//gorm初始化
	core.GormInit()
	// swagger
	core.Swagger()
	//路由初始化
	r := core.RouterInit()
	err := r.Run(":8000")
	if err != nil {
		return
	}

}
