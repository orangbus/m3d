package api_download

import (
	"github.com/gin-gonic/gin"
	"github.com/orangbus/m3d/app/models"
	"github.com/orangbus/m3d/app/response/resp"
	"github.com/orangbus/m3d/pkg/database"
	"github.com/spf13/cast"
	"strings"
)

type ApiDownload struct {
}

func NewApiDownload() *ApiDownload {
	return &ApiDownload{}
}

func (a *ApiDownload) List(ctx *gin.Context) {
	Type := cast.ToInt(ctx.Query("type"))
	var list []models.Download
	var total int64

	db := database.DB.Model(&models.Download{})
	if Type > 0 {
		db = db.Where("type =?", Type)
	}
	db.Count(&total)
	if err := db.Scopes(models.Paginate(ctx)).Find(&list).Error; err != nil {
		resp.Error(ctx, err)
		return
	}
	resp.List(ctx, list, total)
}

func (a *ApiDownload) Store(ctx *gin.Context) {
	var param models.Download
	if err := ctx.ShouldBind(&param); err != nil {
		resp.Error(ctx, err)
		return
	}
	if err := database.DB.Create(&param).Error; err != nil {
		resp.Error(ctx, err)
		return
	}
	resp.Success(ctx, "添加成功")
}

func (a *ApiDownload) Delete(c *gin.Context) {
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
	if err := database.DB.Where("id in (?)", ids).Delete(&models.Download{}).Error; err != nil {
		resp.Error(c, err)
		return
	}
	resp.Success(c, "删除成功")
}

func (a *ApiDownload) Clear(c *gin.Context) {
	if err := database.DB.Where("id > ?", 0).Delete(&models.MovieApi{}).Error; err != nil {
		resp.Error(c, err)
		return
	}
	resp.Success(c, "删除成功")

}

func (a *ApiDownload) Add(ctx *gin.Context) {

}

func (a *ApiDownload) Stop(ctx *gin.Context) {

}
