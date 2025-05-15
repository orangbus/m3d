package api_index

import (
	"github.com/gin-gonic/gin"
	"github.com/orangbus/m3d/app/models"
	"github.com/orangbus/m3d/app/resp"
)

type ApiIndex struct {
}

func NewApiIndex() *ApiIndex {
	return &ApiIndex{}
}

func (a *ApiIndex) Index(ctx *gin.Context) {
	resp.Success(ctx, "success")
	return
}
func (a *ApiIndex) List(c *gin.Context) {
	var list []models.MovieApi

	list = append(list, models.MovieApi{
		Name:   "如意资源",
		Url:    "https://www.ryzyw.com/index.php/vod",
		Type:   0,
		Proxy:  0,
		Status: 0,
	})
	resp.Data(c, list)
}

func (a *ApiIndex) Cate(c *gin.Context) {
	api_url := c.Query("api_url")
	if api_url == "" {
		resp.Error(c, "请选择数据源")
		return
	}
	resp.Data(c, api_url)
}

func (a *ApiIndex) Store(c *gin.Context) {

}

func (a *ApiIndex) Delete(c *gin.Context) {

}
