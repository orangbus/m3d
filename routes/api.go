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

		router.GET("apiList", apiController.List) // 接口列表
		router.GET("cate", apiController.Cate)
		router.POST("store", apiController.Store)
		router.POST("delete", apiController.Delete)

		movieController := api_movie.NewApiMovie()
		router.GET("movie/list", movieController.List)
		router.GET("movie/cate", movieController.Cate)
		router.GET("movie/detail", movieController.Detail)
		router.POST("movie/favorite", movieController.Favorite)
		router.GET("movie/favorite/list", movieController.FavoriteList)
		router.GET("movie/favorite/download", movieController.FavoriteDownload)

		downloadController := api_download.NewApiDownload()
		router.GET("download/list", downloadController.List)
		router.POST("download/store", downloadController.Store)
		router.POST("download/delete", downloadController.Delete)
		router.POST("download/clear", downloadController.Clear)

		router.POST("download/add", downloadController.Add)
		router.GET("download/stop", downloadController.Stop)
	}
}
