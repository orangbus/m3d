package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/orangbus/m3d/app/controller/api/api/api_download"
	"github.com/orangbus/m3d/app/controller/api/api_index"
	"github.com/orangbus/m3d/app/controller/api/api_movie"
	"github.com/orangbus/m3d/app/middleware"
)

func RegisterApiRoutes(r *gin.Engine) {
	router := r.Group("/api/", middleware.ApiAuth())
	{
		apiController := api_index.NewApiIndex()
		router.GET("index", apiController.Index)

		router.GET("list", apiController.List)
		router.GET("cate", apiController.Cate)
		router.POST("store", apiController.Store)
		router.POST("delete", apiController.Delete)

		movieController := api_movie.NewApiMovie()
		router.GET("movie/list", movieController.List)
		router.GET("movie/cate", movieController.Cate)
		router.GET("movie/detail", movieController.Detail)

		downloadController := api_download.NewApiDownload()
		router.GET("download/list", downloadController.List)
		router.POST("download/store", downloadController.Store)
		router.POST("download/delete", downloadController.Delete)
		router.POST("download/clear", downloadController.Clear)

		router.POST("download/add", downloadController.Add)
		router.POST("download/stop", downloadController.Stop)

	}

	//router := r.Group("/api/", middleware.ApiAuth())
	//{
	//apiUser := new(api_user.ApiUserController)
	//router.GET("user", apiUser.UserInfo)       // 用户信息
	//router.POST("user/edit", apiUser.UserEdit) // 编辑
	//
	//// 搜索
	//apiMovie := new(api_movie.ApiMovieController)
	//router.GET("movie/search", apiMovie.Search)                  // 搜索电影
	//router.GET("movie/search/detail", apiMovie.SearchDetail)     // 电影详情
	//router.GET("movie/search/like", apiMovie.SearchLike)         // 相似推荐
	//router.GET("movie/search/analyzer", apiMovie.SearchAnalyzer) // 分词
	//
	//// 历史记录
	//router.POST("movie/recommend", apiMovie.Recommend)
	//router.POST("movie/history", apiMovie.History)
	//router.GET("movie/history/list", apiMovie.HistoryList)
	//router.POST("movie/history/clear", apiMovie.HistoryClear)
	//
	//// 站点授权
	//siteAuth := site_auth.ApiController{}
	//router.GET("apiList", siteAuth.ApiList)               //获取接口列表
	//router.GET("plan", siteAuth.PlanList)                 //获取套餐列表
	//router.GET("notice", siteAuth.NoticeList)             //获取最新公告
	//router.GET("coupon", siteAuth.Coupon)                 // 使用激活码
	//router.POST("email/seed", siteAuth.SeedEmailCode)     // 发送验证码
	//router.POST("email/bind", siteAuth.SeedEmailValidate) // 邮箱认证
	//
	//// 找回账号
	//router.POST("account/forget", siteAuth.Account) // 找回账号
	//}
}
