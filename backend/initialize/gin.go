package initialize

import (
	"bookManage/global"
	"bookManage/router"
	"fmt"
	"github.com/gin-gonic/gin"
)

func InitGin() {
	system := global.Config.System
	r := gin.Default()
	router.SetRouter(r)
	r.Run(fmt.Sprintf("%s:%s", system.Host, system.Port))
}
