package middleware

import (
	"bookManage/schemas/resp"
	"bookManage/util"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")
		if token == "" {
			resp.Fail(c, resp.TokenEmpty)
			c.Abort()
			return
		}
		str := util.RedisUtil.Get(token)
		if str == "" {
			resp.Fail(c, resp.TokenInvalid)
			c.Abort()
			return
		}

		//c.Set("UserId", u.ID)
	}
}
