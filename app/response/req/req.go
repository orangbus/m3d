package req

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func GetPage(c *gin.Context) int {
	page := cast.ToInt(c.DefaultQuery("page", "0"))
	if page < 0 {
		page = 0
	}
	if page > 0 {
		page -= 1
	}
	return page
}
func GetLimit(c *gin.Context) int {
	limit := cast.ToInt(c.DefaultQuery("limit", "10"))
	if limit > 100 {
		limit = 100
	}
	return limit
}
