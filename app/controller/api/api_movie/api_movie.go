package api_movie

import "github.com/gin-gonic/gin"

type ApiMovie struct {
}

func NewApiMovie() *ApiMovie {
	return &ApiMovie{}
}

func (a *ApiMovie) Index(ctx *gin.Context) {

}

func (a *ApiMovie) List(ctx *gin.Context) {

}

func (a *ApiMovie) Store(ctx *gin.Context) {

}

func (a *ApiMovie) Delete(ctx *gin.Context) {

}
