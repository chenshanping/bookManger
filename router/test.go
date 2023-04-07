package router

import "github.com/gin-gonic/gin"

func TestApi(r *gin.Engine) {
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, "test")
	})
}
