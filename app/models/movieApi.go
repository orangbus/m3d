package models

type MovieApi struct {
	TableId
	Name   string `json:"name" form:"name" binding:"required"`
	Url    string `json:"url" form:"url" binding:"required"`
	Type   int    `json:"type" form:"type"`
	Proxy  int    `json:"proxy" form:"proxy"`
	Status int    `json:"status" form:"status"`
	Remark string `json:"remark" form:"remark"`
	TableTime
}
