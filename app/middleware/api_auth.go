package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func ApiAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		hasPrefix := strings.HasPrefix(authorization, "Bearer")
		if !hasPrefix {
			c.AbortWithStatusJSON(419, gin.H{"code": 419, "msg": "请登录"})
		}
		//token := authorization[7:]

		//c.Set("user", user)
		c.Next()
	}
}
