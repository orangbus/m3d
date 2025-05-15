package api_download

import (
	"github.com/gin-gonic/gin"
	"github.com/orangbus/m3d/app/models"
	"github.com/orangbus/m3d/app/resp"
	"github.com/orangbus/m3d/pkg/database"
)

type ApiDownload struct {
}

func NewApiDownload() *ApiDownload {
	return &ApiDownload{}
}

func (a *ApiDownload) List(ctx *gin.Context) {
	var list []models.Download
	var total int64
	if err := database.DB.Find(&list).Error; err != nil {
		resp.Error(ctx, err.Error())
		return
	}
	resp.List(ctx, list, total)
}

func (a *ApiDownload) Store(ctx *gin.Context) {

}

func (a *ApiDownload) Delete(ctx *gin.Context) {

}

func (a *ApiDownload) Clear(ctx *gin.Context) {

}

func (a *ApiDownload) Add(ctx *gin.Context) {

}

func (a *ApiDownload) Stop(ctx *gin.Context) {

}
