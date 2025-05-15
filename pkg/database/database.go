package database

import (
	"database/sql"
	"github.com/orangbus/m3d/pkg/debug"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

var (
	DB    *gorm.DB
	SqlDB *sql.DB
)

/*
*
连接数据库
*/
func Connect(dbConfig gorm.Dialector) {
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		CreateBatchSize: 500,
		NowFunc: func() time.Time {
			return time.Now().UTC() // 设置默认时间为 UTC 时间
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err.Error())
	}
	SqlDB, err = DB.DB()
	SqlDB.SetMaxIdleConns(10)
	SqlDB.SetMaxOpenConns(100)
	SqlDB.SetConnMaxLifetime(time.Hour)
	if err != nil {
		debug.Panic("数据库连接失败", err)
	}
}
