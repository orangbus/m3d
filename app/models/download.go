package models

type Download struct {
	TableId
	ApiId  uint   `json:"api_id" form:"api_id" binding:"required"`
	Name   string `json:"name" form:"name" binding:"required"`
	Url    string `json:"url" form:"url" binding:"required"`
	Status int    `json:"status" form:"status"`
	Proxy  int    `json:"proxy" form:"proxy"`
	Remark string `json:"remark" form:"remark"`
	OutDir string `json:"out_dir" form:"out_dir" binding:"required"` // 保存地址，默认：out .时间，api名称_分类
	TableTime
}
