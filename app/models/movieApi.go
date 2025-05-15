package models

type MovieApi struct {
	TableId
	Name   string `json:"name"`
	Url    string `json:"title"`
	Type   int    `json:"type"`
	Proxy  int    `json:"proxy"`
	Status int    `json:"status"`
	Remark string `json:"remark"`
	TableTime
}
