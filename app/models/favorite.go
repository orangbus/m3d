package models

type Favorite struct {
	TableId
	Name   string `json:"name" form:"name" binding:"required"`
	ApiUrl string `json:"api_url" form:"api_url" binding:"required"`
	TypeId int    `json:"type_id" form:"type_id"`
	TableTime
}
