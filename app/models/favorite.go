package models

type Favorite struct {
	TableId
	Name     string `gorm:"name" json:"name" form:"name" binding:"required"`
	ApiUrl   string `gorm:"api_url" json:"api_url" form:"api_url" binding:"required"`
	TypeId   int    `gorm:"type_id" json:"type_id" form:"type_id" binding:"required"`
	TypeName string `gorm:"type_name" json:"type_name" form:"type_name" binding:"required"`
	TableTime
}

func (f *Favorite) TableName() string {
	return "favorites"
}
