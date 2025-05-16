package models

import (
	"database/sql/driver"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"time"
)

type TableId struct {
	ID uint `gorm:"primaryKey;autoIncrement:true" json:"id"`
}
type TableTime struct {
	CreatedAt Time `gorm:"column:created_at;index;type:datetime" json:"created_at,omitempty"`
	UpdatedAt Time `gorm:"column:updated_at;index;type:datetime" json:"updated_at,omitempty"` // https://github.com/go-gorm/datatypes
}

const timeFormat = "2006-01-02 15:04:05"
const timezone = "Asia/Shanghai"

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormat)
	b = append(b, '"')
	return b, nil
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormat+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormat)
}

func (t Time) local() time.Time {
	loc, _ := time.LoadLocation(timezone)
	return time.Time(t).In(loc)
}

func (t Time) Value() (driver.Value, error) {
	var zeroTime time.Time
	var ti = time.Time(t)
	if ti.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return ti, nil
}

func (t *Time) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = Time(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

// 分页查询： db.Scopes(models.Paginate(c))
func Paginate(r *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := cast.ToInt(r.Query("page"))
		limit := cast.ToInt(r.Query("limit"))
		if page <= 0 {
			page = 1
		}
		if limit <= 0 || limit > 100 {
			limit = 10
		}
		return db.Offset((page - 1) * limit).Limit(limit).Order("id desc")
	}
}
