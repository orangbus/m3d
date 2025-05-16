package api_movie

import (
	"github.com/gin-gonic/gin"
	"github.com/orangbus/m3d/app/response/resp"
	"github.com/orangbus/m3d/pkg/spider"
	"github.com/spf13/cast"
)

type ApiMovie struct {
}

func NewApiMovie() *ApiMovie {
	return &ApiMovie{}
}

func (a *ApiMovie) Index(ctx *gin.Context) {

}

func (a *ApiMovie) List(c *gin.Context) {
	apiUrl := c.Query("api_url")
	TypeId := cast.ToInt(c.Query("type_id"))
	keyword := c.Query("keyword")
	if apiUrl == "" {
		resp.Fail(c, "请选择数据源")
		return
	}
	client := spider.NewSpider(apiUrl)
	if TypeId > 0 {
		client = client.SetTypeId(TypeId)
	}
	if keyword != "" {
		client = client.SetKeyword(keyword)
	}
	res, err := client.Get()
	if err != nil {
		resp.Error(c, err)
		return
	}
	resp.List(c, res.List, res.Total)
}

func (a *ApiMovie) Cate(c *gin.Context) {
	apiUrl := c.Query("api_url")
	if apiUrl == "" {
		resp.Fail(c, "请选择数据源")
		return
	}
	spdier := spider.NewSpider(apiUrl)
	cates, err := spdier.GetMovieCate()
	if err != nil {
		resp.Error(c, err)
		return
	}
	resp.Data(c, cates.Class)
}

func (a *ApiMovie) Detail(ctx *gin.Context) {
	apiUrl := ctx.Query("api_url")
	vodId := ctx.Query("vod_id")
	client := spider.NewSpider(apiUrl)
	res, err := client.GetDetail(vodId)
	if err != nil {
		resp.Error(ctx, err)
		return
	}
	if len(res.List) == 0 {
		resp.Fail(ctx, "未获取到数据")
		return
	}
	resp.Data(ctx, res.List[0])
}
