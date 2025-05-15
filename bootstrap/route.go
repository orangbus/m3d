package bootstrap

import (
	"github.com/gin-gonic/gin"
	"github.com/orangbus/m3d/app/middleware"
	"github.com/orangbus/m3d/pkg/config"
	"github.com/orangbus/m3d/routes"
	"net/http"
	"strings"
)

func SetupRouter(router *gin.Engine) {
	// 注册中间件
	registerGlobalMiddlewares(router)

	// 注册理由
	routes.RegisterWebRoutes(router)
	routes.RegisterApiRoutes(router)
	// 设置接口错误返回
	setup404Handle(router)
}

func registerGlobalMiddlewares(router *gin.Engine) {
	router.Use(
		middleware.Cors(),
		gin.Recovery(),
	)
	if config.GetBool("app.debug") {
		router.Use(gin.Logger())
	}
}

func setup404Handle(router *gin.Engine) {
	router.NoRoute(func(c *gin.Context) {
		acceptString := c.Request.Header.Get("Accept")
		if strings.Contains(acceptString, "text/html") {
			c.String(http.StatusNotFound, "页面不存在")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusNotFound,
				"msg":  "接口不存在",
				"url":  c.Request.RequestURI,
			})
		}
	})
}
