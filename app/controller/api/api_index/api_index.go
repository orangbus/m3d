package api_index

import (
	"github.com/gin-gonic/gin"
	"github.com/orangbus/m3d/app/models"
	"github.com/orangbus/m3d/app/response/resp"
	"github.com/orangbus/m3d/pkg/database"
	"github.com/spf13/cast"
	"strings"
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
	if err := database.DB.Model(&models.MovieApi{}).Find(&list).Error; err != nil {
		resp.Error(c, err)
		return
	}
	resp.Data(c, list)
}

func (a *ApiIndex) Cate(c *gin.Context) {
	api_url := c.Query("api_url")
	if api_url == "" {
		resp.Fail(c, "请选择数据源")
		return
	}
	resp.Data(c, api_url)
}

func (a *ApiIndex) Store(c *gin.Context) {
	var param models.MovieApi
	if err := c.ShouldBind(&param); err != nil {
		resp.Error(c, err)
		return
	}

	var api models.MovieApi
	if err := database.DB.Where("url = ?", param.Url).First(&api).Error; err != nil {
		resp.Error(c, err)
		return
	}
	if api.ID > 0 {
		resp.Fail(c, "当前接口已经存在")
		return
	}

	if err := database.DB.Create(&param).Error; err != nil {
		resp.Error(c, err)
		return
	}
	resp.Data(c, param, "保存成功")
}

func (a *ApiIndex) Delete(c *gin.Context) {
	id := c.Query("id")
	ids := []int{}
	if strings.Contains(id, ",") {
		for _, v := range strings.Split(id, ",") {
			ids = append(ids, cast.ToInt(v))
		}
	} else {
		if cast.ToInt(id) == 0 {
			resp.Fail(c, "请选择要删除的项")
			return
		}
		ids = append(ids, cast.ToInt(id))
	}
	if len(ids) == 0 {
		resp.Fail(c, "请选择要删除的项")
		return
	}
	if err := database.DB.Where("id in (?)", ids).Delete(&models.MovieApi{}).Error; err != nil {
		resp.Error(c, err)
		return
	}
	resp.Success(c, "删除成功")
}
