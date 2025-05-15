package resp

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"strconv"
)

// 成功返回
func Success(c *gin.Context, msg ...string) {
	message := "success"
	if len(msg) > 0 {
		message = msg[0]
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  message,
	})
}

// 错误返回
func Error(c *gin.Context, msg ...string) {
	message := "success"
	if len(msg) > 0 {
		message = msg[0]
	}
	c.JSON(200, gin.H{
		"code": 202,
		"msg":  message,
	})
}

// 数据信息返回
func Data(c *gin.Context, data interface{}, msg ...string) {
	message := "success"
	if len(msg) > 0 {
		message = msg[0]
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
		"msg":  message,
	})
}

// 保存成功
func SaveSuccess(c *gin.Context, msg ...string) {
	message := "保存成功"
	if len(msg) > 0 {
		message = msg[0]
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  message,
	})
}

func UpdateSuccess(c *gin.Context, msg ...string) {
	message := "更新成功"
	if len(msg) > 0 {
		message = msg[0]
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  message,
	})
}

// 删除成功
func DeleteSuccess(c *gin.Context, msg ...string) {
	message := "删除成功"
	if len(msg) > 0 {
		message = msg[0]
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  message,
	})
}

// 列表数据返回
func List(c *gin.Context, data interface{}, total int64, msg ...string) {
	message := "success"
	if len(msg) > 0 {
		message = msg[0]
	}
	c.JSON(200, gin.H{
		"code":  200,
		"data":  data,
		"total": total,
		"msg":   message,
	})
}

func ToInt64(id string) int64 {
	i, _ := strconv.Atoi(id)
	return int64(i)
}

func GetIntVal(c *gin.Context, filed, defaultValue string) int {
	page := c.DefaultQuery(filed, defaultValue)
	return cast.ToInt(page)
}
func GetPage(c *gin.Context, defaultVal string) int {
	page := c.DefaultQuery("page", defaultVal)
	return cast.ToInt(page)
}
func GetLimit(c *gin.Context, defaultVal string) int {
	limit := c.DefaultQuery("page", defaultVal)
	return cast.ToInt(limit)
}
