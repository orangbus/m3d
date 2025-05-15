package models

type Video struct {
	TableId
	Tags []byte `gorm:"column:tags;null;type:json;" json:"tags"`
	TableTime
}
