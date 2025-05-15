package resp

import "github.com/orangbus/m3d/app/models"

type RespMovieVideoList struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	//Page      string          `json:"page"`
	Pagecount int `json:"pagecount"`
	//Limit     string             `json:"limit"`
	Total int64              `json:"total"`
	List  []models.Movies    `json:"list"`
	Class []models.MovieCate `json:"class"`
}

type RespMovieCate struct {
	Class []interface{} `json:"class"`
}
