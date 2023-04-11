package middleware

import (
	"bookManage/global"
	"bookManage/models"
	"bookManage/models/response"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		var u *models.User
		row := global.DB.Where("token=?", token).First(&u).RowsAffected
		if row != 1 {
			response.ResponWrite(c, "当前token错误")
			c.Abort()
		}

		c.Set("UserId", u.ID)
	}
}
