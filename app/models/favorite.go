package models

type Favorite struct {
	TableId
	Name   string `json:"name" from:"name"`
	ApiUrl string `json:"api_url" from:"api_url"`
	TypeId int    `json:"type_id" from:"type_id"`
	TableTime
}
