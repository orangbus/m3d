package models

type Download struct {
	TableId
	ApiId  uint   `json:"api_id"`
	Name   string `json:"name"`
	Url    string `json:"title"`
	Status int    `json:"status"`
	Proxy  int    `json:"proxy"`
	Remark string `json:"remark"`
	TableTime
}
