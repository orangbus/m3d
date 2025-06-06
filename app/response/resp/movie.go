package resp

import "github.com/orangbus/m3d/app/models"

type RespHolidayTour struct {
	Seqid   string `json:"seqid"`
	Code    string `json:"code"`    //"10000"
	Message string `json:"message"` //"SUCCESS"
	Flag    int    `json:"flag"`    // 1
	Data    []RespHolidayTourData
}

type RespHolidayTourData struct {
	Lon        string `json:"lon"`
	Lat        string `json:"lat"`
	ValueRange string `json:"valueRange"`
}

type RespVisitor struct {
	Seqid   string `json:"seqid"`
	Code    string `json:"code"`    //"10000"
	Message string `json:"message"` //"SUCCESS"
	Flag    int    `json:"flag"`    // 1
	Data    []RespVisitorData
}

type RespVisitorData struct {
	ProvName    string `json:"provName"`    // "上海市",
	ProvCode    string `json:"ProvCode"`    // "310000",
	CityCode    string `json:"CityCode"`    // "310100",
	RegionName  string `json:"RegionName"`  // "",
	DataType    string `json:"DataType"`    // "102",
	RegionCode  string `json:"RegionCode"`  // "",
	AreaCode    string `json:"AreaCode"`    // "",
	SrcAreaName string `json:"SrcAreaName"` //  "三亚市",
	PersonTimes string `json:"PersonTimes"` // "298374",
	CityName    string `json:"CityName"`    // "上海市",
	SrcAreaCode string `json:"SrcAreaCode"` // "460200",
	AreaName    string `json:"AreaName"`    // "",
	AreaType    string `json:"AreaType"`    // "2"
}

type RespMovieVideoList struct {
	Code      int                `json:"code"`
	Msg       string             `json:"msg"`
	Page      any                `json:"page"`
	Pagecount int                `json:"pagecount"`
	Limit     any                `json:"limit"`
	Total     int64              `json:"total"`
	List      []models.Movies    `json:"list"`
	Class     []models.MovieCate `json:"class"`
}

type RespMovieCate struct {
	Class []interface{} `json:"class"`
}
