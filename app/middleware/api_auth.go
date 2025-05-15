package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/orangbus/m3d/pkg/config"
	"strings"
)

func ApiAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorization := c.GetHeader("Authorization")
		hasPrefix := strings.HasPrefix(authorization, "Bearer")
		if !hasPrefix {
			c.AbortWithStatusJSON(419, gin.H{"code": 419, "msg": "请输入授权秘钥"})
		}
		token := authorization[7:]
		if token != config.GetString("app.key") {
			c.AbortWithStatusJSON(419, gin.H{"code": 419, "msg": "秘钥错误"})
		}
		c.Next()
	}
}
