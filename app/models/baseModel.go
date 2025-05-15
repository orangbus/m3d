package models

import (
	"time"
)

type LocalTime time.Time

type TableId struct {
	ID uint `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
}

type TableTime struct {
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"index"`
}

// 格式化时间
func (t *LocalTime) MarshalJSON() ([]byte, error) {
	formatted := time.Time(*t).Format("2006-01-02 15:04:05")
	return []byte(formatted), nil
}

func (t *LocalTime) UnmarshalJSON(data []byte) (err error) {
	now, _ := time.ParseInLocation("2006-01-02 15:04:05", string(data), time.Local)
	*t = LocalTime(now)
	return
}
