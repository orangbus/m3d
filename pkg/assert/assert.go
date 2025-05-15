package assert

import (
	"github.com/gin-gonic/gin"
	"github.com/orangbus/m3d/pkg/console"
	"log"
)

func Error(name string, err error) bool {
	if err != nil {
		console.Error(name, err)
		return false
	}
	return false
}

func HasError(c *gin.Context, msg string, err error) bool {
	if err != nil {
		log.Println(err.Error())
		//resp.Error(c, msg)
		return true
	}
	return false
}

func IsTrue(c *gin.Context, condition bool, msg string) bool {
	if condition {
		//resp.Error(c, msg)
		return true
	}
	return false
}

func IsEmpty(c *gin.Context, value string, msg string) bool {
	if value == "" {
		//resp.Error(c, msg)
		return true
	}
	return false
}
