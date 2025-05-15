package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/orangbus/m3d/pkg/config"
)

func RegisterWebRoutes(router *gin.Engine) {
	router.Static("/assets", "ui/assets")
	router.Static("/images", "ui/images")
	router.LoadHTMLGlob("ui/index.html")
	router.Static("upload", "./upload")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": config.GetString("app.name"),
		})
	})

	// 接口登录
	//apiLogin := new(api_auth.LoginController)
	//router.POST("/api/login", apiLogin.Login)       // 前端 + 接口 登录
	//router.POST("/api/register", apiLogin.Register) // 前端 + 接口 登录
	//router.POST("/api/email", apiLogin.SeedEmail)   // 前端 + 接口 发送邮件
	//router.GET("/api/website", apiLogin.Website)    // 前端 + 接口 登录
	//
	//adminLogin := new(admin_auth.LoginController)
	//router.POST("/admin/login", adminLogin.Login)
	//router.GET("/admin/website", adminLogin.Website)
	//router.GET("/admin/log", adminLogin.Website)
	//adminMovie := admin_movie.NewWebSocketController()
	//router.GET("movie/logs", adminMovie.Conn) // 实时日志
}
