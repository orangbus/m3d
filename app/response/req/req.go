package req

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func GetPageLimit(c *gin.Context) (int, int) {
	page := cast.ToInt(c.DefaultQuery("page", "0"))
	limit := cast.ToInt(c.DefaultQuery("limit", "10"))
	if page < 0 {
		page = 0
	}
	if page > 0 {
		page -= 1
	}
	if limit > 100 {
		limit = 100
	}
	return page, limit
}
